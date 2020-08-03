// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

import (
	"net/http"
	"time"

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
	state   string
	born    time.Time
	err     error
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	m := gobot.NewMaster()
	m.AutoRun = true
	api := api.NewAPI(m)
	r := &Robot{
		Master: m,
		api: api,
		state: "Init",
		born: time.Now(),
	}
	r.addCommands(r.Master)
	r.Master.AddRobot(platform.NewRobot())
	return r
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

func (m *Robot) CurrentState(s string) {
	m.state = s
}
