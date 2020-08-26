// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package driver implements the munbot gobot.Driver interface.
package driver

import (
	"fmt"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

var _ Driver = &Munbot{}

type Driver interface {
	gobot.Driver
}

type Munbot struct {
	gobot.Driver
	name string
	conn gobot.Connection
	gobot.Eventer
}

func New(a adaptor.Adaptor) *Munbot {
	return &Munbot{
		name:    "munbot",
		conn:    a,
		Eventer: a.Eventer(),
	}
}

// gobot interface

func (d *Munbot) Name() string {
	return d.name
}

func (d *Munbot) SetName(name string) {
	d.name = name
}

func (d *Munbot) Connection() gobot.Connection {
	return d.conn
}

func (d *Munbot) Start() error {
	log.Printf("Start %s driver.", d.name)
	// user login handler
	if err := d.On(event.UserLogin, func(data interface{}) {
		if data != nil {
			log.Debug("user login")
			sess := data.(event.Session)
			ev := fmt.Sprintf("%s.%s", event.UserLogin, sess.Sid)
			d.Publish(ev, event.Error{})
		}
	}); err != nil {
		log.Panic(err)
	}
	// user logout handler
	if err := d.On(event.UserLogout, func(data interface{}) {
		if data != nil {
			log.Print("USER LOGOUT")
		}
	}); err != nil {
		log.Panic(err)
	}
	return nil
}

func (d *Munbot) Halt() error {
	log.Printf("Halt %s driver.", d.name)
	return nil
}
