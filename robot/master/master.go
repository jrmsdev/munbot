// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

import (
	"net/http"
	"sync"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	"github.com/munbot/master/env"
	"github.com/munbot/master/internal/api/wapp"
	"github.com/munbot/master/log"
	"github.com/munbot/master/platform"
)

type Config struct {
	Name string
}

var _ Munbot = &Robot{}

type Robot struct {
	*gobot.Master
	name  string
	api   wapp.Api
	state string
	born  time.Time
	err   error
	exitc chan<- bool
	stop  chan bool
	rw    *sync.RWMutex
}

func New() Munbot {
	return NewRobot()
}

func NewRobot() *Robot {
	m := gobot.NewMaster()
	m.AutoRun = false
	r := &Robot{
		Master: m,
		name:   env.Get("MUNBOT"),
		api:    wapp.New(api.NewAPI(m)),
		state:  "Init",
		born:   time.Now(),
		stop:   make(chan bool, 0),
		rw:     new(sync.RWMutex),
	}
	r.addCommands(r.Master)
	r.Master.Start()
	platform.AddRobots(r.Master)
	return r
}

// gobot interface

func (m *Robot) Start() error {
	log.Debugf("start master robot %s...", m.name)
	autorun := false
	return m.Master.Robots().Start(autorun)
}

func (m *Robot) Stop() error {
	log.Debugf("stop master robot %s...", m.name)
	return m.Master.Robots().Stop()
}

// munbot interface

func (m *Robot) CurrentState(s string) {
	m.state = s
}

func (m *Robot) ExitNotify(c chan<- bool) {
	m.exitc = c
}

func (m *Robot) Configure(c *Config, wc *wapp.Config) error {
	m.rw.Lock()
	defer m.rw.Unlock()
	if c.Name != "" {
		m.name = c.Name
	}
	m.api.Configure(wc)
	return nil
}

func (m *Robot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.api.ServeHTTP(w, r)
}
