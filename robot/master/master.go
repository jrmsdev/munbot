// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

import (
	"gobot.io/x/gobot"
)

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Master
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	return &Robot{Master: gobot.NewMaster()}
}
