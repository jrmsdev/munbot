// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master implements the master robot.
package master

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/log"
)

// Robot works around a gobot.Master.
type Robot struct {
	*gobot.Master
}

// New creates a new master robot.
func New() *Robot {
	return &Robot{Master: gobot.NewMaster()}
}

// Start starts the core runtime and then the gobot master robot.
func (m *Robot) Start() error {
	log.Print("Start master robot.")
	m.Master.AutoRun = true
	return m.Master.Start()
}

// Stop stops the gobot master robot and then the core runtime.
func (m *Robot) Stop() error {
	return m.Master.Stop()
}

// Run runs the robot's main loop.
func (m *Robot) Run() error {
	defer m.Stop()
	return m.Start()
}
