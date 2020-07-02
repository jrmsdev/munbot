// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/internal/config"
)

type Master struct {
	*config.Section
	Name  *config.StringValue `json:"name,omitempty"`
	Hostname  *config.StringValue `json:"hostname,omitempty"`
	Api   *Api                `json:"api,omitempty"`
	Robot *Robot              `json:"robot,omitempty"`
}

func NewMaster(m *config.Manager) *Master {
	s := m.NewSection("master")
	return &Master{
		Section: s,
		Name:    s.NewString("name", flags.Name),
		Hostname:    s.NewString("hostname", "localhost"),
		Api:     newApi(s, true),
		Robot:   newRobot(s, flags.Name, true),
	}
}

type Api struct {
	Enable *config.BoolValue     `json:"enable,omitempty"`
	Addr   *config.StringValue   `json:"addr,omitempty"`
	Port   *config.IntValue      `json:"port,omitempty"`
	Cert   *config.FilepathValue `json:"cert,omitempty"`
	Key    *config.FilepathValue `json:"key,omitempty"`
	Path   *config.PathValue     `json:"path,omitempty"`
}

func newApi(s *config.Section, enable bool) *Api {
	return &Api{
		Enable: s.NewBool("api.enable", enable),
		Addr:   s.NewString("api.addr", "0.0.0.0"),
		Port:   s.NewInt("api.port", 3000),
		Cert:   s.NewFilepath("api.cert", filepath.FromSlash("ssl/api/cert.pem")),
		Key:    s.NewFilepath("api.key", filepath.FromSlash("ssl/api/key.pem")),
		Path:   s.NewPath("api.path", "/api"),
	}
}

type Robot struct {
	Enable  *config.BoolValue   `json:"enable,omitempty"`
	Name    *config.StringValue `json:"name,omitempty"`
	AutoRun *config.BoolValue   `json:"autorun,omitempty"`
}

func newRobot(s *config.Section, name string, enable bool) *Robot {
	return &Robot{
		Enable:  s.NewBool("robot.enable", enable),
		Name:    s.NewString("robot.name", name),
		AutoRun: s.NewBool("robot.autorun", true),
	}
}
