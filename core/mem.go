// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/config"
)

type Mem struct {
	Flags    *Flags
	Cfg      *config.Config
	CfgFlags *config.Flags
}

func newMem() *Mem {
	return &Mem{}
}
