// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot implements the core munbot robot.
package robot

import (
	"os"
	"os/signal"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/driver"
	"github.com/munbot/master/v0/driver/api"
	"github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

// Munbot implements the core worker robot.
type Munbot struct {
	*gobot.Robot
	conn    adaptor.Adaptor
	AutoRun bool
	gobot.Eventer
}

func New(conn adaptor.Adaptor) *Munbot {
	r := &Munbot{
		conn:    conn,
		AutoRun: true,
		Eventer: conn.Eventer(),
	}
	r.Robot = gobot.NewRobot(
		"munbot",
		[]gobot.Connection{r.conn},
		[]gobot.Device{
			driver.New(r.conn),
			api.NewDriver(r.conn),
		},
		r.Work,
	)
	r.Robot.Eventer = r.Eventer
	return r
}

// Gobot returns the internal gobot reference, to be added to a master gobot.
func (r *Munbot) Gobot() *gobot.Robot {
	r.Robot.AutoRun = r.AutoRun
	return r.Robot
}

func (r *Munbot) Work() {
	log.Debug("start work...")
	r.AddEvent(event.Fail)
	if err := r.Once(event.Fail, func(data interface{}) {
		if data != nil {
			log.Info("Failure!")
		}
	}); err != nil {
		log.Panic(err)
	}
	r.Publish(event.ApiStart, nil)
	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Info("OS interrupt!")
	r.Publish(event.Fail, nil)
}
