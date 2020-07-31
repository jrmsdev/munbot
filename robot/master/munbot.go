// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/robot/worker"
)

type Munbot interface {
	AddWorker(*gobot.Robot) worker.Munbot
	Start() error
	Stop() error
}
