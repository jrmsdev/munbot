// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"github.com/munbot/master/config"
)

type Server interface {
	Configure(cfg *config.Section) error
	Start() error
	Stop() error
}
