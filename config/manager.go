// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Manager struct {
	*registry
}

func New() *Manager {
	return &Manager{newReg()}
}

func (m *Manager) NewSection(name string) *Section {
	s := newSection(name)
	m.registry.sect.PushBack(s)
	return s
}

func (m *Manager) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	m.registry.db[name] = v
	return v
}

func (m *Manager) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	m.registry.db[name] = v
	return v
}

func (m *Manager) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	m.registry.db[name] = v
	return v
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
