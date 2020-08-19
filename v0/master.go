// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master.
package master

import (
	"github.com/munbot/master/v0/version"
)

// Robot works around a gobot.Master.
type Robot struct {
}

// New creates a new master robot.
func New() *Robot {
	return &Robot{}
}

// Version returns version information.
func (m *Robot) Version() *version.Info {
	return new(version.Info)
}

// Run runs the robot's main loop.
func (m *Robot) Run() error {
	return nil
}
