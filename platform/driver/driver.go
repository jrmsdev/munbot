// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package driver implements the munbot gobot.Driver interface.
package driver

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/munbot/master/platform/adaptor"
)

var _ Munbot = &Driver{}

type Munbot interface {
	gobot.Driver
}

const Hello string = "hello"

type Driver struct {
	gobot.Eventer
	gobot.Commander
	name       string
	connection gobot.Connection
	interval   time.Duration
	halt       chan bool
}

func New(a adaptor.Munbot) *Driver {
	m := &Driver{
		name:       "Munbot",
		connection: a,
		interval:   500 * time.Millisecond,
		halt:       make(chan bool, 0),
		Eventer:    gobot.NewEventer(),
		Commander:  gobot.NewCommander(),
	}

	m.AddEvent(Hello)

	m.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		return m.Hello()
	})

	return m
}

func (m *Driver) Connection() gobot.Connection {
	return m.connection
}

func (m *Driver) Name() string { return m.name }

func (m *Driver) SetName(name string) { m.name = name }

func (m *Driver) Start() error {
	go func() {
		for {
			m.Publish(m.Event(Hello), m.Hello())

			select {
			case <-time.After(m.interval):
			case <-m.halt:
				return
			}
		}
	}()
	return nil
}

func (m *Driver) Halt() error {
	m.halt <- true
	return nil
}

func (m *Driver) adaptor() adaptor.Munbot {
	return m.Connection().(adaptor.Munbot)
}

func (m *Driver) Hello() string {
	return "hello from " + m.Name() + "!"
}

func (m *Driver) Ping() string {
	return m.adaptor().Ping()
}
