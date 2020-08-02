// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api defines and implements master api server interface.
package api

import (
	"fmt"
	"net/http"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

var _ Server = &Api{}

type Api struct {
	enable bool
	server *http.Server
}

func New() Server {
	return &Api{server: new(http.Server)}
}

func (a *Api) Configure(kfl *flags.Flags, cfg *config.Section) error {
	a.enable = cfg.GetBool("enable")
	a.server.Addr = fmt.Sprintf("%s:%d", kfl.ApiAddr, kfl.ApiPort)
	return nil
}

func (a *Api) Start() error {
	return nil
}

func (a *Api) Stop() error {
	return nil
}
