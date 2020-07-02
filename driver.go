// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/jrmsdev/munbot/log"
)

const Hello string = "hello"

type Driver struct {
	name       string
	connection gobot.Connection
	interval   time.Duration
	halt       chan bool
	gobot.Eventer
	gobot.Commander
}

func NewDriver(a *Adaptor) *Driver {
	d := &Driver{
		name:       a.Name(),
		connection: a,
		interval:   500 * time.Millisecond,
		halt:       make(chan bool, 0),
		Eventer:    gobot.NewEventer(),
		Commander:  gobot.NewCommander(),
	}
	d.AddEvent(Hello)
	d.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		return d.Hello()
	})
	return d
}

// gobot interface methods

func (d *Driver) Name() string { return d.name }

func (d *Driver) SetName(name string) { d.name = name }

func (d *Driver) Start() error {
	log.Print("Start driver ", d.name, "...")
	go func() {
		for {
			d.Publish(d.Event(Hello), d.Hello())
			select {
			case <-time.After(d.interval):
			case <-d.halt:
				return
			}
		}
	}()
	return nil
}

func (d *Driver) Halt() error {
	log.Print("Halt driver ", d.name, "...")
	d.halt <- true
	return nil
}

func (d *Driver) Connection() gobot.Connection {
	return d.connection
}

// custom methods

func (d *Driver) adaptor() *Adaptor {
	return d.Connection().(*Adaptor)
}

func (d *Driver) Hello() string {
	return "hello from " + d.Name() + "!"
}

func (d *Driver) Ping() string {
	return d.adaptor().Ping()
}
