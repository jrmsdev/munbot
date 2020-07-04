// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/internal/config"
)

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
