// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"container/list"
	"encoding/json"
	"io"
	"io/ioutil"
)

type Manager struct {
	*registry
	sect *list.List
}

func New() *Manager {
	return &Manager{newReg(), list.New()}
}

func (m *Manager) NewSection(name string) *Section {
	s := newSection(name)
	m.sect.PushBack(s)
	return s
}

func (m *Manager) Dump() {
	for e := m.sect.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Section)
		s.Dump()
	}
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
	fh.Write(blob)
	fh.Write([]byte("\n"))
	return nil
}
