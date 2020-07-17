// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"strconv"
)

type Section struct {
	name string
	m Map
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
		// TODO: debug log about missing option
		return ""
	}
	return v
}

func (s *Section) GetBool(name string) bool {
	r, err := strconv.ParseBool(s.Get(name))
	if err != nil {
		// TODO: log about parsing error
		return false
	}
	return r
}
