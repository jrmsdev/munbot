// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package platform implements internal gobot platform.
package platform

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/robot/worker"
)

func NewRobot() *gobot.Robot {
	return worker.New().Gobot()
}
