// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"strconv"

	"github.com/munbot/master/log"
)

type Section struct {
	name string
	m    Map
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) HasOption(name string) bool {
	_, found := s.m[name]
	return found
}

func (s *Section) Get(name string) string {
	v, found := s.m[name]
	if !found {
		log.Debugf("config missing option: %s.%s", s.name, name)
		return ""
	}
	return v
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
