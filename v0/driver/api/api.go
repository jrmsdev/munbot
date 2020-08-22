// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api implements munbot platform's api driver.
package api

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/internal/core"
	"github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

type Driver struct {
	conn adaptor.Adaptor
	name string
	srv  core.ApiServer
	gobot.Eventer
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		conn: a,
		name: "munbot.api",
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
	a.conn.GobotApi()
	a.srv = core.NewApiServer()
	if err := a.srv.Configure(); err != nil {
		return log.Errorf("Api server configure: %v", err)
	}
	a.On(event.ApiStart, func(data interface{}) {
		log.Print("Start api server.")
	})
	a.On(event.ApiStop, func(data interface{}) {
		log.Print("Stop api server.")
	})
	return nil
}

func (a *Driver) Halt() error {
	log.Printf("Halt %s driver.", a.name)
	a.srv = nil
	return nil
}
