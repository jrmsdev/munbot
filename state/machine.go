// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"time"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

var _ Machine = &SM{}

type Machine interface {
	Config() *config.Config
	ConfigFlags() *config.Flags
	CoreFlags() *core.Flags
	Runtime() core.Runtime
	Init(*config.Flags, *core.Flags) error
	Run(context.Context) error
	SetState(StateID)
}

type hist struct {
	Time  time.Time `json:"time"`
	State string    `json:"state"`
}

// SM implements the state.Machine interface.
type SM struct {
	hist      []*hist
	newst     bool
	st        State
	init      State
	configure State
	start     State
	cfg       *config.Config
	cfgFlags  *config.Flags
	coreFlags *core.Flags
	rt        core.Runtime
}

func NewMachine() Machine {
	m := &SM{}
	m.rt = core.NewRuntime()
	m.hist = make([]*hist, 0)
	m.init = NewInitState(m)
	m.configure = newConfigure(m)
	m.start = newStart(m)
	m.SetState(Init)
	return m
}

func (m *SM) SetState(sid StateID) {
	var s State
	switch sid {
	case Init:
		s = m.init
	case Configure:
		s = m.configure
	case Start:
		s = m.start
	}
	log.Debugf("%v set %s", m.st, s)
	m.st = s
	m.hist = append(m.hist, &hist{time.Now(), s.String()})
	m.newst = true
}

func (m *SM) Config() *config.Config {
	return m.cfg
}

func (m *SM) ConfigFlags() *config.Flags {
	return m.cfgFlags
}

func (m *SM) CoreFlags() *core.Flags {
	return m.coreFlags
}

func (m *SM) Runtime() core.Runtime {
	return m.rt
}

func (m *SM) Init(cf *config.Flags, fl *core.Flags) error {
	m.cfg = config.New()
	m.cfgFlags = cf
	m.coreFlags = fl
	return nil
}

func (m *SM) Run(ctx context.Context) error {
	var err error
	rc := OK
	for rc == OK {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			rc = DONE
		default:
			if m.newst {
				m.newst = false
				n := m.st.String()
				log.Debugf("%s run", n)
				ctx, rc = m.st.Run(ctx)
				log.Debugf("%s status %s", n, stMap[rc])
			} else {
				log.Debug("no new state to run... exit!")
				rc = EXIT
			}
		}
	}
	log.Debugf("exit status %s", stMap[rc])
	if rc == ERROR {
		err = m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return err
}
