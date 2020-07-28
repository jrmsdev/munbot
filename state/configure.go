// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

var ConfigureError = errors.New("internal error: run Master.Configure first")

type ConfigureState struct {
	m   *Machine
	err error
}

func newConfigure(m *Machine) *ConfigureState {
	return &ConfigureState{m: m}
}

func (s *ConfigureState) String() string {
	return "Configure"
}

func (s *ConfigureState) Error() error {
	return s.err
}

func (s *ConfigureState) Run(ctx context.Context) Status {
	select {
	case <-ctx.Done():
		return DONE
	default:
		if err := s.configure(); err != nil {
			s.err = err
			return ERROR
		}
	}
	return EXIT
}

func (s *ConfigureState) configure() error {
	s.m.Config.SetDefaults(config.Defaults)
	if err := s.m.Config.Load(s.m.ConfigFlags.Profile); err != nil {
		return log.Error(s.err)
	}
	s.m.CoreFlags.Parse(s.m.Config)
	return s.m.Runtime.Configure(s.m.Config, s.m.ConfigFlags, s.m.CoreFlags)
}
