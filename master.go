// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core lib for the munbot project.
package master

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/version"
)

// Version returns the running version information.
func Version() *version.Info {
	return new(version.Info)
}

type Master struct {
	rt core.Runtime
}

func New() *Master {
	return NewMaster(core.NewRuntime())
}

func NewMaster(rt core.Runtime) *Master {
	return &Master{rt: rt}
}

func (m *Master) Init(ctx context.Context) (context.Context, error) {
	return m.rt.Init(ctx)
}

func (m *Master) Configure(kf *core.Flags, cf *config.Flags, c *config.Config) error {
	return m.rt.Configure(kf, cf, c)
}

func (m *Master) Start() error {
	return m.rt.Start()
}

func (m *Master) Stop() error {
	return m.rt.Stop()
}
