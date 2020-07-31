// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/config"
)

type State interface {
	Init() error
	Configure(*Flags, *config.Flags, *config.Config) error
	Start() error
	Stop() error
}
