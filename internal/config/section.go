// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"path"
)

type Section struct {
	opt  *list.List
	idx  map[string]Value
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

func (s *Section) Dump(out io.Writer, listAll bool, section, opt string) {
	for e := s.opt.Front(); e != nil; e = e.Next() {
		v := e.Value.(Value)
		f := s.filter(section, opt, s.name, v.Name())
		if (listAll || v.modified() || section != "") && f {
			io.WriteString(out,
				fmt.Sprintf("%s.%s=%s\n", s.name, v.Name(), v.String()))
		}
	}
}

func (s *Section) filter(sect, opt, xs, xn string) bool {
	if sect == "" || sect == xs {
		if opt == "" || opt == xn {
			return true
		}
	}
	return false
}

func (s *Section) register(name string, v Value) {
	s.idx[name] = v
	s.opt.PushBack(v)
}

func (s *Section) Update(opt, newval string) error {
	if _, ok := s.idx[opt]; !ok {
		return errors.New(fmt.Sprintf("invalid config section '%s' option: '%s'", s.name, opt))
	}
	return s.idx[opt].Update(newval)
}

func (s *Section) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	s.register(name, v)
	return v
}

func (s *Section) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	s.register(name, v)
	return v
}

func (s *Section) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	s.register(name, v)
	return v
}

func (s *Section) NewFilepath(name string, defval string, forceAbs bool) *FilepathValue {
	v := &FilepathValue{newValue("filepath", name), defval, forceAbs}
	s.register(name, v)
	return v
}

func (s *Section) NewPath(name string, defval string) *PathValue {
	forceAbs := path.IsAbs(defval)
	v := &PathValue{newValue("filepath", name), defval, forceAbs}
	s.register(name, v)
	return v
}
