// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/internal/config"
)

type Master struct {
	*config.Section
	Name *config.StringValue `json:"name,omitempty"`
	Api *Api `json:"api,omitempty"`
	Robot *Robot `json:"robot,omitempty"`
}

func NewMaster(m *config.Manager) *Master {
	s := m.NewSection("master")
	return &Master{
		Section: s,
		Name:    s.NewString("name", "munbot"),
		Api:     newApi(s, true),
		Robot:   newRobot(s, "munbot", true),
	}
}

type Api struct {
	Enable *config.BoolValue `json:"enable,omitempty"`
	Host *config.StringValue `json:"host,omitempty"`
	Port *config.IntValue `json:"port,omitempty"`
}

func newApi(s *config.Section, enable bool) *Api {
	return &Api{
		Enable: s.NewBool("api.enable", enable),
		Host: s.NewString("api.host", ""),
		Port: s.NewInt("api.port", 3000),
	}
}

type Robot struct {
	Name *config.StringValue `json:"name,omitempty"`
	AutoRun *config.BoolValue `json:"autorun,omitempty"`
}

func newRobot(s *config.Section, name string, autoRun bool) *Robot {
	return &Robot{
		Name:    s.NewString("robot.name", name),
		AutoRun: s.NewBool("robot.autorun", autoRun),
	}
}
