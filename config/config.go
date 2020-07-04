// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/internal/config"
)

type Master struct {
	*config.Section
	Name     *config.StringValue `json:"name,omitempty"`
	Hostname *config.StringValue `json:"hostname,omitempty"`
	Api      *Api                `json:"api,omitempty"`
	Robot    *Robot              `json:"robot,omitempty"`
}

func NewMaster(m *config.Manager) *Master {
	s := m.NewSection("master")
	return &Master{
		Section:  s,
		Name:     s.NewString("name", flags.Name),
		Hostname: s.NewString("hostname", "localhost"),
		Api:      newApi(s, true),
		Robot:    newRobot(s, flags.Name, true),
	}
}
