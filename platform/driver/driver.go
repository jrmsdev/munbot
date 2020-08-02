// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package driver implements the munbot gobot.Driver interface.
package driver

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/munbot/master/platform/adaptor"
)

var _ Driver = &Munbot{}

type Driver interface {
	gobot.Driver
}

const Hello string = "hello"

type Munbot struct {
	gobot.Eventer
	gobot.Commander
	name       string
	connection gobot.Connection
	interval   time.Duration
	halt       chan bool
}

func New(a adaptor.Adaptor) *Munbot {
	m := &Munbot{
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

func (m *Munbot) Connection() gobot.Connection {
	return m.connection
}

func (m *Munbot) Name() string { return m.name }

func (m *Munbot) SetName(name string) { m.name = name }

func (m *Munbot) Start() error {
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

func (m *Munbot) Halt() error {
	m.halt <- true
	return nil
}

func (m *Munbot) adaptor() adaptor.Adaptor {
	return m.Connection().(adaptor.Adaptor)
}

func (m *Munbot) Hello() string {
	return "hello from " + m.Name() + "!"
}

func (m *Munbot) Ping() string {
	return m.adaptor().Ping()
}
