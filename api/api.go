// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api defines and implements master api server interface.
package api

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

var _ Server = &Api{}

type Api struct {
}

func New() Server {
	return &Api{}
}

func (a *Api) Configure(kfl *flags.Flags, cfg *config.Section) error {
	return nil
}

func (a *Api) Start() error {
	return nil
}

func (a *Api) Stop() error {
	return nil
}
