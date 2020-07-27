// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import "errors"

var ConfigureError = errors.New("internal error: run Master.Configure first")

type ConfigureState struct {
	m   *Machine
	err error
}

func newConfigure(m *Machine) *ConfigureState {
	return &ConfigureState{m: m}
}

func (s *ConfigureState) Error() error {
	return s.err
}

func (s *ConfigureState) Run() Status {
	if s.m.Config == nil {
		s.err = ConfigureError
		return PANIC
	}
	s.m.setState(s.m.init)
	return OK
}
