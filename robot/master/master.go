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

	"github.com/munbot/master/api/wapp"
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
		name:   "master",
		api:    wapp.New(api.NewAPI(m)),
		state:  "Init",
		born:   time.Now(),
		stop:   make(chan bool, 1),
		rw:     new(sync.RWMutex),
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
	log.Printf("Uptime %s", time.Since(m.born))
	return nil
}

func (m *Robot) Stop() error {
	var err error
	for _, bot := range *m.Master.Robots() {
		log.Debugf("robot %s stop...", bot.Name)
		if err = bot.Stop(); err != nil {
			log.Warnf("stop robot %s error: %v", bot.Name, err)
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
