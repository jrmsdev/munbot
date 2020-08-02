// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

import (
	"net/http"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Master
	api *api.API
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	m := gobot.NewMaster()
	return &Robot{Master: m, api: api.NewAPI(m)}
}

func (m *Robot) Configure(kfl *flags.Flags, cfl *config.Flags, cfg *config.Config) error {
	return nil
}

func (m *Robot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.api.ServeHTTP(w, r)
}
