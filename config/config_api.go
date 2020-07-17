// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/munbot/master/config/parser"
)

type Api struct {
	Enable bool
	Addr   string
	Port   int
}

func (c *Api) load(s *parser.Section) {
	c.Enable = s.GetBool("enable")
	c.Addr = s.Get("addr")
	c.Port = s.GetInt("port")
}
