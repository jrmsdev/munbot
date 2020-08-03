// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"net/http"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

type Munbot interface {
	Configure(kfl *flags.Flags, cfl *config.Flags, cfg *config.Config) error
	Start() error
	Stop() error
	Running() bool
	ServeHTTP(http.ResponseWriter, *http.Request)
	CurrentState(string)
}
