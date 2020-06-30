// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Manager struct {
	registry
}

func New() *Manager {
	return &Manager{make(registry)}
}

func (c *Manager) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	c.registry[name] = v
	return v
}

func (c *Manager) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	c.registry[name] = v
	return v
}

func (c *Manager) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	c.registry[name] = v
	return v
}
