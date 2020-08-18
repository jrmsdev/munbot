// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package platform implements internal gobot platform.
package platform

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/robot/worker"
)

func AddRobots(m *gobot.Master) {
	m.AddRobot(worker.New().Gobot())
}
