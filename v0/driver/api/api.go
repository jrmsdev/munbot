// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api implements munbot platform's api driver.
package api

import (
	"sync"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/core"
	"github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

type Driver struct {
	*sync.Mutex
	conn adaptor.Adaptor
	name string
	srv  core.ApiServer
	wg   *sync.WaitGroup
	err  error
	gobot.Eventer
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		Mutex:   new(sync.Mutex),
		conn:    a,
		name:    "munbot.api",
		wg:      new(sync.WaitGroup),
		Eventer: a.Eventer(),
	}
}

// gobot interface

func (a *Driver) Name() string {
	return a.name
}

func (a *Driver) SetName(name string) {
	a.Lock()
	defer a.Unlock()
	a.name = name
}

func (a *Driver) Connection() gobot.Connection {
	return a.conn
}

func (a *Driver) Start() error {
	log.Printf("Start %s driver.", a.name)
	a.Lock()
	defer a.Unlock()
	if a.srv != nil {
		return log.Error("Api server alread started.")
	}
	a.err = nil
	a.srv = core.NewApiServer()
	// configure
	log.Print("Configure api server.")
	if err := a.srv.Configure(); err != nil {
		return log.Errorf("Api server configure: %v", err)
	}
	if env.GetBool("MBAPI") {
		ga := a.conn.GobotApi()
		if env.GetBool("MBAPI_DEBUG") {
			ga.Debug()
		}
		ga.AddRobeauxRoutes()
		a.srv.Mount(env.Get("MBAPI_PATH"), ga)
	}
	// stop handler
	a.wg.Add(1)
	if err := a.Once(event.ApiStop, func(data interface{}) {
		defer a.wg.Done()
		log.Debugf("got %s: %v", event.ApiStop, data)
		if run := data.(bool); !run {
			return
		}
		log.Debug("stop api server")
		if err := a.srv.Stop(); err != nil {
			a.err = log.Error(err)
		}
	}); err != nil {
		return log.Error(err)
	}
	// start handler
	a.wg.Add(1)
	if err := a.Once(event.ApiStart, func(data interface{}) {
		defer a.wg.Done()
		log.Debugf("got %s: %v", event.ApiStart, data)
		if run := data.(bool); !run {
			return
		}
		log.Debug("start api server")
		if err := a.srv.Start(); err != nil {
			log.Error(err)
			a.Publish(event.Fail, event.Error{event.ApiStart, err})
		}
	}); err != nil {
		return log.Error(err)
	}
	log.Debug("start done")
	return nil
}

func (a *Driver) Halt() error {
	log.Printf("Halt %s driver.", a.name)
	a.Lock()
	defer a.Unlock()
	log.Debugf("publish %s", event.ApiStop)
	a.Publish(event.ApiStop, true)
	log.Debug("wait...")
	a.wg.Wait()
	a.srv = nil
	return a.err
}
