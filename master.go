// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core lib for the munbot project.
package master

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/state"
	"github.com/munbot/master/version"
)

// Version returns the running version information.
func Version() *version.Info {
	return new(version.Info)
}

type Master struct {
	sm *state.Machine
}

func New() *Master {
	return &Master{sm: state.NewMachine()}
}

func (m *Master) Run() error {
	return m.sm.Run(context.Background())
}

func (m *Master) Configure(cf *config.Flags, fl *Flags) error {
	cfg := config.New()
	if err := cfg.Load(cf.Profile); err != nil {
		return err
	}
	m.sm.Config = cfg
	return nil
}
