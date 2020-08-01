// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/log"
)

var _ State = &SHalt{}

type SHalt struct {
	m  Machine
	rt *Mem
}

func newHalt(m Machine, rt *Mem) State {
	return &SHalt{m: m, rt: rt}
}

func (s *SHalt) Init() error {
	return ErrInit
}

func (s *SHalt) Configure() error {
	return ErrConfigure
}

func (s *SHalt) Start() error {
	return ErrStart
}

func (s *SHalt) Run() error {
	return ErrRun
}

func (s *SHalt) Stop() error {
	return ErrStop
}

func (s *SHalt) Halt() error {
	log.Print("Halt")
	return nil
}
