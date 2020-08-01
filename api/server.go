// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

type Server interface {
	Configure(kfl *flags.Flags, cfg *config.Section) error
	Start() error
	Stop() error
}
