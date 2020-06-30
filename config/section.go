// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
	"errors"
	"fmt"
	"io"
)

type Section struct {
	opt *list.List
	idx map[string]Value
	name string
}

func newSection(name string) *Section {
	return &Section{list.New(), make(map[string]Value), name}
}

func (s *Section) String() string {
	return s.name
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) Dump(out io.Writer, listAll bool) {
	for e := s.opt.Front(); e != nil; e = e.Next() {
		v := e.Value.(Value)
		if listAll || v.modified() {
			io.WriteString(out,
				fmt.Sprintf("%s.%s=%s\n", s.name, v.Name(), v.String()))
		}
	}
}

func (s *Section) Update(opt, newval string) error {
	if _, ok := s.idx[opt]; !ok {
		return errors.New(fmt.Sprintf("invalid config section %s option: %s", s.name, opt))
	}
	return s.idx[opt].Update(newval)
}

func (s *Section) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	s.opt.PushBack(v)
	s.idx[name] = v
	return v
}

func (s *Section) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	s.opt.PushBack(v)
	s.idx[name] = v
	return v
}

func (s *Section) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	s.opt.PushBack(v)
	s.idx[name] = v
	return v
}
