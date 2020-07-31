// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

//~ import (
	//~ "github.com/munbot/master/config"
//~ )

var _ State = &SInit{}

type SInit struct {
	m Machine
	rt *Mem
}

func NewInit(m Machine, rt *Mem) State {
	return &SInit{m: m, rt: rt}
}

func (s *SInit) Init() error {
	s.m.SetState(Configure)
	return nil
}

func (s *SInit) Configure() error {
	return ErrConfigure
}

func (s *SInit) Start() error {
	return ErrStart
}

func (s *SInit) Stop() error {
	return ErrStop
}
