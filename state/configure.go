// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

var _ State = &ConfigureState{}

var ConfigureError = errors.New("internal error: run Master.Configure first")

type ConfigureState struct {
	m   *SM
	id  StateID
	err error
}

func newConfigure(m *SM) *ConfigureState {
	return &ConfigureState{m: m, id: Configure}
}

func (s *ConfigureState) String() string {
	return s.id.String()
}

func (s *ConfigureState) Error() error {
	return s.err
}

func (s *ConfigureState) Run(ctx context.Context) (context.Context, Status) {
	select {
	case <-ctx.Done():
		return ctx, DONE
	default:
		if err := s.configure(); err != nil {
			s.err = err
			return ctx, ERROR
		}
	}
	if err := s.m.SetState(Start); err != nil {
		s.err = log.Errorf("%s: %s", err, Start)
		return ctx, ERROR
	}
	return ctx, OK
}

func (s *ConfigureState) configure() error {
	cfg := s.m.Config()
	cfg.SetDefaults(config.Defaults)
	cfl := s.m.ConfigFlags()
	if err := cfg.Load(cfl.Profile); err != nil {
		return log.Error(s.err)
	}
	kfl := s.m.CoreFlags()
	kfl.Parse(cfg)
	return s.m.Runtime().Configure(cfg, cfl, kfl)
}
