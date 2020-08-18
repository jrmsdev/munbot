// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"net/http"
	"time"

	"github.com/munbot/master/internal/api/wapp"
)

type Munbot interface {
	Configure(*Config, *wapp.Config) error
	Start() error
	Stop() error
	Running() bool
	ServeHTTP(http.ResponseWriter, *http.Request)
	CurrentState(string)
	ExitNotify(chan<- bool)
	Uptime() time.Duration
}
