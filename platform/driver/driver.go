// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package driver implements the munbot gobot.Driver interface.
package driver

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/munbot/master/log"
	"github.com/munbot/master/platform/adaptor"
)

var _ Driver = &Munbot{}

type Driver interface {
	gobot.Driver
}

type Munbot struct {
	gobot.Eventer
	gobot.Commander
	name string
	conn gobot.Connection
}

func New(a adaptor.Adaptor) *Munbot {
	return &Munbot{
		name:      "munbot",
		conn:      a,
		Eventer:   gobot.NewEventer(),
		Commander: gobot.NewCommander(),
	}
}

// gobot interface

func (m *Munbot) Connection() gobot.Connection {
	return m.conn
}

func (m *Munbot) Name() string {
	return m.name
}

func (m *Munbot) SetName(name string) {
	m.name = name
}

func (m *Munbot) Start() error {
	log.Debug("start")
	return nil
}

func (m *Munbot) Halt() error {
	log.Debug("halt")
	return nil
}

// munbot interface

func (m *Munbot) adaptor() adaptor.Adaptor {
	return m.Connection().(adaptor.Adaptor)
}

func (m *Munbot) interval() time.Duration {
	return m.adaptor().Interval()
}
