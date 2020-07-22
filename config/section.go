// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"

	"github.com/munbot/master/config/internal/parser"
	"github.com/munbot/master/log"
)

type Section struct {
	name string
	h    *parser.Config
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) HasOption(name string) bool {
	return s.h.HasOption(s.name, name)
}

func (s *Section) Get(name string) string {
	return s.h.Get(s.name, name)
}

func (s *Section) GetBool(name string) bool {
	r, err := strconv.ParseBool(s.Get(name))
	if err != nil {
		log.Errorf("config option %s.%s parse error: %s", s.name, name, err)
		return false
	}
	return r
}

func (s *Section) GetInt(name string) int {
	r, err := strconv.Atoi(s.Get(name))
	if err != nil {
		log.Errorf("config option %s.%s parse error: %s", s.name, name, err)
		return 0
	}
	return r
}
