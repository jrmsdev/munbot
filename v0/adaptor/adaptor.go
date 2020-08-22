// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package adaptor implements the munbot gobot.Adaptor interface.
package adaptor

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	"github.com/munbot/master/v0/internal/core"
	"github.com/munbot/master/v0/internal/event"
	"github.com/munbot/master/v0/log"
)

var _ gobot.Connection = &Munbot{}

type Adaptor interface {
	gobot.Adaptor
	Interval() time.Duration
	SetInterval(time.Duration)
	GobotApi() *api.API
	Eventer() gobot.Eventer
}

type Munbot struct {
	Adaptor
	master   *gobot.Master
	name     string
	interval time.Duration
	evtr     event.Manager
}

func New(m *gobot.Master) *Munbot {
	return &Munbot{
		master:   m,
		name:     "munbot",
		interval: 300 * time.Millisecond,
		evtr:     event.NewManager(),
	}
}

// gobot interface

func (m *Munbot) Name() string {
	return m.name
}

func (m *Munbot) SetName(name string) {
	m.name = name
}

func (m *Munbot) Connect() error {
	log.Printf("Connect %s platform.", m.name)
	log.Debug("lock core runtime")
	core.Lock()
	return nil
}

func (m *Munbot) Finalize() error {
	log.Printf("Finalize %s platform.", m.name)
	log.Debug("wait eventer to finish...")
	m.evtr.Wait()
	log.Debug("unlock core runtime")
	core.Unlock()
	return nil
}

// munbot interface

func (m *Munbot) Interval() time.Duration {
	return m.interval
}

func (m *Munbot) SetInterval(d time.Duration) {
	m.interval = d
}

func (m *Munbot) GobotApi() *api.API {
	return api.NewAPI(m.master)
}

func (m *Munbot) Eventer() gobot.Eventer {
	return m.evtr
}
