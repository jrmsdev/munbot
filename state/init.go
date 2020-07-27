// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import "errors"

var InitError = errors.New("init: run Configure first")

type InitState struct {
	m   *Machine
	err error
}

func newInit(m *Machine) *InitState {
	return &InitState{m: m}
}

func (s *InitState) String() string {
	return "InitState"
}

func (s *InitState) Error() error {
	return s.err
}

func (s *InitState) Run() Status {
	if s.m.Config == nil {
		s.err = InitError
		return ERROR
	}
	return EXIT
}
