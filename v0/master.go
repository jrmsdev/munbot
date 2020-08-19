// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master implements the master robot.
package master

import (
	"gobot.io/x/gobot"
)

// Robot works around a gobot.Master.
type Robot struct {
	*gobot.Master
}

// New creates a new master robot.
func New() *Robot {
	return &Robot{Master: gobot.NewMaster()}
}

// Run runs the robot's main loop.
func (m *Robot) Run() error {
	return nil
}
