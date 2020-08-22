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
	conn adaptor.Adaptor
	name string
	srv  core.ApiServer
	wg   *sync.WaitGroup
	gobot.Eventer
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		conn: a,
		name: "munbot.api",
		wg:   new(sync.WaitGroup),
		Eventer: gobot.NewEventer(),
	}
}

// gobot interface

func (a *Driver) Name() string {
	return a.name
}

func (a *Driver) SetName(name string) {
	a.name = name
}

func (a *Driver) Connection() gobot.Connection {
	return a.conn
}

func (a *Driver) Start() error {
	log.Printf("Start %s driver.", a.name)
	a.srv = core.NewApiServer()
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
	a.wg.Add(1)
	if err := a.Once(event.ApiStart, func(data interface{}) {
		log.Print("Start api server.")
		defer a.wg.Done()
		a.srv.Start()
	}); err != nil {
		return log.Error(err)
	}
	a.wg.Add(1)
	if err := a.Once(event.ApiStop, func(data interface{}) {
		log.Print("Stop api server.")
		defer a.wg.Done()
		a.srv.Stop()
	}); err != nil {
		return log.Error(err)
	}
	a.Publish(event.ApiStart, nil)
	return nil
}

func (a *Driver) Halt() error {
	log.Printf("Halt %s driver.", a.name)
	a.Publish(event.ApiStop, nil)
	a.wg.Wait()
	a.srv = nil
	return nil
}
