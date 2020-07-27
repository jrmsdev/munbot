// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
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

func (m *Machine) Run() error {
	rc := OK
	for rc == OK {
		n := m.st.String()
		log.Debugf("%s run", n)
		rc = m.st.Run()
		log.Debugf("%s status %s", n, stMap[rc])
	}
	if rc == ERROR {
		return m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return nil
}
