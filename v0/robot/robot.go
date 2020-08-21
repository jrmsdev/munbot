// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot implements the core munbot robot.
package robot

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/driver"
	"github.com/munbot/master/v0/driver/api"
	"github.com/munbot/master/v0/log"
)

// Munbot implements the core worker robot.
type Munbot struct {
	*gobot.Robot
}

func New() *Munbot {
	r := &Munbot{}
	conn := adaptor.New()
	r.Robot = gobot.NewRobot(
		"munbot",
		[]gobot.Connection{conn},
		[]gobot.Device{
			driver.New(conn),
			api.NewDriver(conn),
		},
		r.Work,
	)
	return r
}

// Gobot returns the internal gobot reference, to be added to a master gobot.
func (r *Munbot) Gobot() *gobot.Robot {
	return r.Robot
}

func (r *Munbot) Work() {
	log.Debug("start work...")
}
