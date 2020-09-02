// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package sshd implements munbot platform's ssh server driver.
package sshd

import (
	"sync"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/internal/core"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/net"
)

type Driver struct {
	*sync.Mutex
	conn adaptor.Adaptor
	name string
	srv  core.SSHServer
	wg   *sync.WaitGroup
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		Mutex: new(sync.Mutex),
		conn:  a,
		name:  "munbot.sshd",
		wg:    new(sync.WaitGroup),
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
	// start
	s.wg.Add(1)
	go func() {
		log.Debug("start ssh server")
		defer s.wg.Done()
		if err := s.srv.Start(); err != nil {
			log.Panic(err)
		}
	}()
	return nil
}

func (s *Driver) Halt() error {
	log.Printf("Halt %s driver.", s.name)
	s.Lock()
	defer s.Unlock()
	log.Debug("stop ssh server")
	err := s.srv.Stop()
	if err != nil {
		log.Error(err)
	}
	log.Debug("wait...")
	s.wg.Wait()
	if err == nil {
		s.srv = nil
	}
	return err
}

// munbot interface

func (s *Driver) Addr() *net.Addr {
	return s.srv.Addr()
}
