// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"gobot.io/x/gobot"
)

type Munbot interface {
	Gobot() *gobot.Robot
}
