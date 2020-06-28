// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"gobot.io/x/gobot"
)

type Master struct {
	*gobot.Master
}

func New() *Master {
	return &Master{gobot.NewMaster()}
}
