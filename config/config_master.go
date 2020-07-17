// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/munbot/master/config/parser"
)

type Master struct {
	Name   string `json:"name,omitempty"`
	Enable bool   `json:"enable,omitempty"`
}

func (c *Master) load(s *parser.Section) {
	c.Name = s.Get("name")
	c.Enable = s.GetBool("enable")
}
