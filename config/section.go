// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
)

type Section struct {
	opt *list.List
	name string
}

func newSection(name string) *Section {
	return &Section{list.New(), name}
}

func (s *Section) String() string {
	return s.name
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	s.opt.PushBack(v)
	return v
}
