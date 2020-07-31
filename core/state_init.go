// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

var _ State = &SInit{}

type SInit struct {
	m  Machine
	rt *Mem
}

func newInit(m Machine, rt *Mem) State {
	return &SInit{m: m, rt: rt}
}

func (s *SInit) Init() error {
	return nil
}

func (s *SInit) Configure() error {
	return s.m.SetState(Run)
}

func (s *SInit) Start() error {
	return ErrStart
}

func (s *SInit) Stop() error {
	return ErrStop
}