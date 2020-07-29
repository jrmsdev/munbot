// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

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
	State() StateID
	SetState(StateID) error
}

// SM implements the state.Machine interface.
type SM struct {
	newst     bool
	stid      StateID
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
	m.init = NewInitState(m)
	m.configure = newConfigure(m)
	m.start = newStart(m)
	m.SetState(Init)
	return m
}

func (m *SM) State() StateID {
	return m.stid
}

var ErrSetSameState error = errors.New("sm: set same state")
var ErrSetInvalid error = errors.New("sm: set invalid state")

func (m *SM) SetState(stid StateID) error {
	var s State
	if stid == m.stid {
		return ErrSetSameState
	}
	switch stid {
	case Init:
		s = m.init
	case Configure:
		s = m.configure
	case Start:
		s = m.start
	default:
		return ErrSetInvalid
	}
	log.Debugf("%v set %s", m.st, s)
	m.st = s
	m.stid = stid
	m.newst = true
	return nil
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
				log.Debugf("%s status %s", n, rc)
			} else {
				log.Debug("no new state to run... exit!")
				rc = EXIT
			}
		}
	}
	log.Debug(rc)
	if rc == ERROR {
		err = m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return err
}
