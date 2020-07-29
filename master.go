// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core lib for the munbot project.
package master

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/state"
	"github.com/munbot/master/version"
)

// Version returns the running version information.
func Version() *version.Info {
	return new(version.Info)
}

type Master struct {
	rt core.Runtime
	sm state.Machine
}

func New() *Master {
	rt := core.NewRuntime()
	return &Master{rt: rt, sm: state.NewMachine(rt)}
}

func (m *Master) Init(cf *config.Flags, fl *core.Flags) error {
	return m.sm.Init(cf, fl)
}

func (m *Master) Run() error {
	return m.sm.Run(context.Background())
}
