// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

var _ State = &SRun{}

type SRun struct {
	m  Machine
	rt *Mem
}

func newRun(m Machine, rt *Mem) State {
	return &SRun{m: m, rt: rt}
}

func (s *SRun) Init() error {
	return ErrInit
}

func (s *SRun) Configure() error {
	return ErrConfigure
}

func (s *SRun) Start() error {
	return nil
}

func (s *SRun) Stop() error {
	return ErrStop
}
