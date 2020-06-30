// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Manager struct {
	sect *list.List
	idx map[string]*Section
}

func New() *Manager {
	return &Manager{list.New(), make(map[string]*Section)}
}

func (m *Manager) NewSection(name string) *Section {
	s := newSection(name)
	m.sect.PushBack(s)
	m.idx[name] = s
	return s
}

func (m *Manager) Dump(out io.Writer, listAll bool, filter string) {
	for e := m.sect.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Section)
		section, opt := m.filter(filter)
		s.Dump(out, listAll, section, opt)
	}
}

func (m *Manager) filter(f string) (string, string) {
	s := ""
	n := ""
	i := strings.Split(f, ".")
	s = i[0]
	if len(i) >= 1 {
		n = strings.Join(i[1:], ".")
	}
	return s, n
}

func (m *Manager) Update(section, opt, newval string) error {
	if _, ok := m.idx[section]; !ok {
		return errors.New(fmt.Sprintf("invalid config section: %s", section))
	}
	return m.idx[section].Update(opt, newval)
}

func (m *Manager) Read(obj interface{}, fh io.Reader) error {
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(blob, obj); err != nil {
		return err
	}
	return nil
}

func (m *Manager) Write(obj interface{}, fh io.Writer) error {
	blob, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	if _, err := fh.Write(blob); err != nil {
		return err
	}
	return nil
}
