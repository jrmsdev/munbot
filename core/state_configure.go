// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

//~ import (
//~ 	"github.com/munbot/master/config"
//~ )

var _ State = &SConfigure{}

type SConfigure struct {
	m  Machine
	rt *Mem
}

func newConfigure(m Machine, rt *Mem) State {
	return &SConfigure{m: m, rt: rt}
}

func (s *SConfigure) Init() error {
	return ErrInit
}

func (s *SConfigure) Configure() error {
	return s.m.SetState(Run)
}

func (s *SConfigure) Start() error {
	return ErrStart
}

func (s *SConfigure) Stop() error {
	return ErrStop
}
