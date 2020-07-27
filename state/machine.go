// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

type Machine struct {
	configure State
	init      State
	st        State
	Config    *config.Config
}

func NewMachine() *Machine {
	m := &Machine{}
	m.configure = newConfigure(m)
	m.init = newInit(m)
	m.setState(m.configure)
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
