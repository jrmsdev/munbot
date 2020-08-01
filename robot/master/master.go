// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/api"
	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Master
	api api.Server
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	return &Robot{
		Master: gobot.NewMaster(),
		api: api.New(),
	}
}

func (m *Robot) Configure(kfl *flags.Flags, cfl *config.Flags) error {
	return m.api.Configure(kfl)
}
