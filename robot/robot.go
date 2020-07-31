// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot binds both master and worker robot interfaces.
package robot

import (
	"github.com/munbot/master/robot/master"
	"github.com/munbot/master/robot/worker"
)

func NewMaster() master.Munbot {
	return master.New()
}

func NewBot() worker.Munbot {
	return worker.New()
}
