// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package worker defines and implements the worker robot interface.
package worker

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/platform/adaptor"
	"github.com/munbot/master/platform/driver"
)

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Robot
	adaptor adaptor.Adaptor
	driver  driver.Driver
}

func New() Munbot {
	r := &Robot{
		adaptor: adaptor.New(),
	}
	r.driver = driver.New(r.adaptor)
	r.Robot = gobot.NewRobot(
		"Munbot",
		[]gobot.Connection{r.adaptor},
		[]gobot.Device{r.driver},
		r.work,
	)
	return r
}

// munbot interface

func (r *Robot) Gobot() *gobot.Robot {
	return r.Robot
}

// work

func (r *Robot) work() {
}
