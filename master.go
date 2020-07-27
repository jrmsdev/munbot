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

func (m *Master) Init(cf *config.Flags, fl *Flags) error {
	m.sm.Config = config.New()
	m.sm.ConfigFlags = cf
	return nil
}

func (m *Master) Run() error {
	return m.sm.Run(context.Background())
}
