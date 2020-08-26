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
	evtr     event.Eventer
}

func New(m *gobot.Master) *Munbot {
	return &Munbot{
		master:   m,
		name:     "munbot",
		interval: 300 * time.Millisecond,
		evtr:     event.NewEventer(),
	}
}

// gobot interface

func (a *Munbot) Name() string {
	return a.name
}

func (a *Munbot) SetName(name string) {
	a.name = name
}

func (a *Munbot) Connect() error {
	log.Printf("Connect %s platform.", a.name)
	log.Debug("lock core runtime")
	core.Lock()
	return nil
}

func (a *Munbot) Finalize() error {
	log.Printf("Finalize %s platform.", a.name)
	a.evtr.Publish(event.Fail, nil)
	log.Debug("wait eventer to finish...")
	a.evtr.Wait()
	log.Debug("unlock core runtime")
	core.Unlock()
	return nil
}

// munbot interface

func (a *Munbot) Interval() time.Duration {
	return a.interval
}

func (a *Munbot) SetInterval(d time.Duration) {
	a.interval = d
}

func (a *Munbot) GobotApi() *api.API {
	return api.NewAPI(a.master)
}

func (a *Munbot) Eventer() gobot.Eventer {
	return a.evtr
}
