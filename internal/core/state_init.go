// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/api"
	"github.com/munbot/master/api/wapp"
	"github.com/munbot/master/config"
	"github.com/munbot/master/console"
	"github.com/munbot/master/env"
	"github.com/munbot/master/internal/auth"
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
		s.rt.Console = console.New()
	}
	return nil
}

func (s *SInit) Configure() error {
	log.Print("Configure...")
	cfg := s.m.Config()
	cfl := s.m.ConfigFlags()
	cfg.SetDefaults(config.Defaults)
	if err := cfg.Load(); err != nil {
		return log.Error(err)
	}

	log.Print("Configure auth manager...")
	cadir := cfl.Profile.GetPath("auth")
	if err := s.rt.Auth.Configure(cadir); err != nil {
		return err
	}

	log.Print("Configure master robot...")
	mcfg := &master.Config{
		Name: env.Get("MUNBOT"),
	}
	wappcfg := &wapp.Config{
		Enable: env.GetBool("MBAPI"),
		Debug:  env.GetBool("MBAPI_DEBUG"),
		Path:   env.Get("MBAPI_PATH"),
	}
	if err := s.rt.Master.Configure(mcfg, wappcfg); err != nil {
		return log.Error(err)
	}

	log.Print("Configure master api...")
	apiEnable := env.GetBool("MBAPI")
	apiCfg := &api.ServerConfig{
		Enable: apiEnable,
		Addr:   env.Get("MBAPI_ADDR"),
		Port:   env.GetUint("MBAPI_PORT"),
	}
	if err := s.rt.Api.Configure(apiCfg); err != nil {
		return log.Error(err)
	}
	if apiEnable {
		s.rt.Api.Mount(env.Get("MBAPI_PATH"), s.rt.Master)
	}

	log.Print("Configure master console...")
	consCfg := &console.Config{
		Enable: env.GetBool("MBCONSOLE"),
		Addr:   env.Get("MBCONSOLE_ADDR"),
		Port:   env.GetUint("MBCONSOLE_PORT"),
		Auth:   s.rt.Auth,
	}
	if err := s.rt.Console.Configure(consCfg); err != nil {
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
