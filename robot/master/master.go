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
	"github.com/munbot/master/platform"
)

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Master
	api     *api.API
	cfginit bool
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	m := gobot.NewMaster()
	m.AddRobot(platform.NewRobot())
	api := api.NewAPI(m)
	return &Robot{Master: m, api: api}
}

func (m *Robot) Configure(kfl *flags.Flags, cfl *config.Flags, cfg *config.Config) error {
	if !m.cfginit {
		if kfl.ApiEnable {
			m.api.AddRobeauxRoutes()
		}
		m.cfginit = true
	}
	if kfl.ApiDebug {
		m.api.Debug()
	}
	return nil
}

func (m *Robot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.api.ServeHTTP(w, r)
}
