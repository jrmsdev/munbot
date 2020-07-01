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
}

func NewMaster(m *config.Manager) *Master {
	s := m.NewSection("master")
	return &Master{
		Section: s,
		Name:    s.NewString("name", "munbot"),
		Api:     newApi(s),
	}
}

type Api struct {
	Enable *config.BoolValue `json:"enable,omitempty"`
}

func newApi(s *config.Section) *Api {
	return &Api{
		Enable: s.NewBool("api.enable", false),
	}
}
