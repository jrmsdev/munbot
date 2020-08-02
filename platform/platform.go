// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package platform implements internal gobot platform.
package platform

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/platform/adaptor"
	"github.com/munbot/master/platform/driver"
	"github.com/munbot/master/robot/worker"
)

type Adaptor adaptor.Adaptor

func NewAdaptor() Adaptor {
	return adaptor.New()
}

type Driver driver.Driver

func NewDriver(a Adaptor) Driver {
	return driver.New(a)
}

func NewRobot() *gobot.Robot {
	return worker.New().Gobot()
}
