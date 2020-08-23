// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package sshd implements munbot platform's ssh server driver.
package sshd

import (
	"sync"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	//~ "github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/core"
	//~ "github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

type Driver struct {
	conn adaptor.Adaptor
	name string
	srv  core.SSHServer
	wg   *sync.WaitGroup
	gobot.Eventer
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		conn:    a,
		name:    "munbot.sshd",
		wg:      new(sync.WaitGroup),
		Eventer: a.Eventer(),
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
	// configure
	// stop handler
	// start handler
	return nil
}

func (a *Driver) Halt() error {
	log.Printf("Halt %s driver.", a.name)
	return nil
}
