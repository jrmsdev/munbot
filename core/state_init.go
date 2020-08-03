// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/api"
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
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
	log.Print("Init...")
	if s.rt.Master == nil {
		s.rt.Master = master.New()
		s.rt.Api = api.New()
	}
	return nil
}

func (s *SInit) Configure() error {
	log.Print("Configure...")
	cfg := s.m.Config()
	cfl := s.m.ConfigFlags()
	kfl := s.m.CoreFlags()
	cfg.SetDefaults(config.Defaults)
	if err := cfg.Load(cfl.Profile); err != nil {
		return log.Error(err)
	}
	kfl.Parse(cfg)
	if err := s.rt.Master.Configure(kfl, cfl, cfg); err != nil {
		return log.Error(err)
	}
	apiCfg := &api.ServerConfig{
		Enable: kfl.ApiEnable,
		Addr:   kfl.ApiAddr,
		Port:   kfl.ApiPort,
	}
	if err := s.rt.Api.Configure(apiCfg); err != nil {
		return log.Error(err)
	}
	if kfl.ApiEnable {
		s.rt.Api.Mount("/", s.rt.Master)
	}
	return s.m.SetState(Run)
}

func (s *SInit) Start() error {
	return ErrStart
}

func (s *SInit) Run() error {
	return ErrRun
}

func (s *SInit) Stop() error {
	return ErrStop
}

func (s *SInit) Halt() error {
	return ErrHalt
}
