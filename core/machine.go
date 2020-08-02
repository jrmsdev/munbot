// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

type Machine interface {
	SetState(StateID) error
	Config() *config.Config
	ConfigFlags() *config.Flags
	CoreFlags() *flags.Flags
}
