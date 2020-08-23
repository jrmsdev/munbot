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
	*sync.Mutex
	conn adaptor.Adaptor
	name string
	srv  core.SSHServer
	wg   *sync.WaitGroup
	gobot.Eventer
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		Mutex:   new(sync.Mutex),
		conn:    a,
		name:    "munbot.sshd",
		wg:      new(sync.WaitGroup),
		Eventer: a.Eventer(),
	}
}

// gobot interface

func (s *Driver) Name() string {
	return s.name
}

func (s *Driver) SetName(name string) {
	s.Lock()
	defer s.Unlock()
	s.name = name
}

func (s *Driver) Connection() gobot.Connection {
	return s.conn
}

func (s *Driver) Start() error {
	log.Printf("Start %s driver.", s.name)
	s.Lock()
	defer s.Unlock()
	if s.srv != nil {
		return log.Error("SSH server already started")
	}
	s.srv = core.NewSSHServer()
	// configure
	log.Printf("Configure ssh server.")
	if err := s.srv.Configure(); err != nil {
		return log.Error(err)
	}
	// stop handler
	// start handler
	return nil
}

func (s *Driver) Halt() error {
	log.Printf("Halt %s driver.", s.name)
	return nil
}
