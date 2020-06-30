// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Manager struct {
}

func New() *Manager {
	return &Manager{}
}

func (c *Manager) NewString(name string, defval string) *StringValue {
	return &StringValue{newValue("string", name), defval}
}

func (c *Manager) NewInt(name string, defval int) *IntValue {
	return &IntValue{newValue("int", name), defval}
}

func (c *Manager) NewBool(name string, defval bool) *BoolValue {
	return &BoolValue{newValue("bool", name), defval}
}
