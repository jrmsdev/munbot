// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/api"
	"github.com/munbot/master/auth"
	"github.com/munbot/master/config"
	"github.com/munbot/master/console"
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
		log.Print("Init auth manager...")
		s.rt.Auth = auth.New()
		log.Print("Init master robot...")
		s.rt.Master = master.New()
		log.Print("Init master api...")
		s.rt.Api = api.New()
		log.Print("Init master console...")
		s.rt.Console = console.NewServer()
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
	log.Print("Configure auth manager...")
	cadir := cfl.Profile.GetPath("auth")
	if err := s.rt.Auth.Configure(cadir); err != nil {
		return log.Error(err)
	}
	log.Print("Configure master robot...")
	if err := s.rt.Master.Configure(kfl, cfl, cfg); err != nil {
		return log.Error(err)
	}
	log.Print("Configure master api...")
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
	log.Print("Configure master console...")
	if err := console.Configure(s.rt.Console, kfl, s.rt.Auth); err != nil {
		return log.Error(err)
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
