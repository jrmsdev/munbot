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
	"github.com/munbot/master/log"
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
	exitc   chan<- bool
	stop    chan bool
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	m := gobot.NewMaster()
	m.AutoRun = false
	api := api.NewAPI(m)
	r := &Robot{
		Master: m,
		api: api,
		state: "Init",
		born: time.Now(),
		stop: make(chan bool, 1),
	}
	r.addCommands(r.Master)
	r.Master.Start()
	r.Master.AddRobot(platform.NewRobot())
	return r
}

// gobot interface

func (m *Robot) Start() error {
	for _, bot := range *m.Master.Robots() {
		autorun := false
		log.Debugf("robot %s start...", bot.Name)
		if err := bot.Start(autorun); err != nil {
			m.Stop()
			return err
		}
	}
	<-m.stop
	return nil
}

func (m *Robot) Stop() error {
	var err error
	for _, bot := range *m.Master.Robots() {
		log.Debugf("robot %s stop...", bot.Name)
		if err = bot.Stop(); err != nil {
			log.Warn("stop robot %s error: %v", bot.Name, err)
		}
	}
	defer close(m.stop)
	m.stop <- true
	return err
}

// munbot interface

func (m *Robot) CurrentState(s string) {
	m.state = s
}

func (m *Robot) ExitNotify(c chan<- bool) {
	m.exitc = c
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
