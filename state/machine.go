// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

type Machine struct {
	st          State
	init        State
	configure   State
	Config      *config.Config
	ConfigFlags *config.Flags
	CoreFlags   *core.Flags
	Runtime     *core.Runtime
}

func NewMachine() *Machine {
	m := &Machine{}
	m.init = newInit(m)
	m.configure = newConfigure(m)
	m.st = m.init
	return m
}

func (m *Machine) setState(s State) {
	log.Debugf("%v set %s", m.st, s)
	m.st = s
}

func (m *Machine) Run(ctx context.Context) error {
	var err error
	rc := OK
	for rc == OK {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			rc = DONE
		default:
			n := m.st.String()
			log.Debugf("%s run", n)
			rc = m.st.Run(ctx)
			log.Debugf("%s status %s", n, stMap[rc])
		}
	}
	if rc == ERROR {
		err = m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return err
}
