// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// master state machine.
package state

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

type Status int

const (
	OK Status = iota
	EXIT
	ERROR
	PANIC
)

type State interface {
	Error() error
	Run() Status
}

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
	log.Debugf("set %T", s)
	m.st = s
}

func (m *Machine) Run() error {
	rc := OK
	for rc == OK {
		log.Debugf("run %T", m.st)
		rc = m.st.Run()
	}
	if rc == ERROR {
		return m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return nil
}
