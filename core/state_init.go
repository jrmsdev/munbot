// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/robot/master"
)

var _ State = &SInit{}

type SInit struct {
	m  Machine
	rt *Mem
}

func newInit(m Machine, rt *Mem) State {
	return &SInit{m: m, rt: rt}
}

func (s *SInit) Init() error {
	s.rt.Master = master.New()
	return nil
}

func (s *SInit) Configure() error {
	cfg := s.m.Config()
	cfl := s.m.ConfigFlags()
	kfl := s.m.CoreFlags()
	cfg.SetDefaults(config.Defaults)
	if err := cfg.Load(cfl.Profile); err != nil {
		return err
	}
	kfl.Parse(cfg)
	s.rt.Cfg = cfg
	s.rt.CfgFlags = cfl
	s.rt.CoreFlags = kfl
	return s.m.SetState(Run)
}

func (s *SInit) Start() error {
	return ErrStart
}

func (s *SInit) Stop() error {
	return ErrStop
}
