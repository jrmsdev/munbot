// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot implements the core munbot robot.
package robot

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/driver/api"
	"github.com/munbot/master/v0/driver/sshd"
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
			api.NewDriver(r.conn),
			sshd.NewDriver(r.conn),
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
	log.Debug("no work! =)")
}
