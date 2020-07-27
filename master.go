// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core lib for the munbot project.
package master

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/version"
)

// Version returns the running version information.
func Version() *version.Info {
	return new(version.Info)
}

type Master struct {
}

func New() *Master {
	return &Master{}
}

func (m *Master) Configure(cf *config.Flags, fl *Flags) error {
	return nil
}

func (m *Master) Run() error {
	return nil
}
