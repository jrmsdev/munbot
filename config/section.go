// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
	"fmt"
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

func (s *Section) Dump() {
	for e := s.opt.Front(); e != nil; e = e.Next() {
		k := e.Value.(Value)
		fmt.Printf("%s.%s=\n", s.name, k.Name())
	}
}

func (s *Section) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	s.opt.PushBack(v)
	return v
}

func (s *Section) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	s.opt.PushBack(v)
	return v
}

func (s *Section) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	s.opt.PushBack(v)
	return v
}
