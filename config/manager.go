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

func (c *Manager) NewString(name string, defval string) *StringValue {
	v := &StringValue{newValue("string", name), defval}
	c.registry.db[name] = v
	return v
}

func (c *Manager) NewInt(name string, defval int) *IntValue {
	v := &IntValue{newValue("int", name), defval}
	c.registry.db[name] = v
	return v
}

func (c *Manager) NewBool(name string, defval bool) *BoolValue {
	v := &BoolValue{newValue("bool", name), defval}
	c.registry.db[name] = v
	return v
}

func (c *Manager) Read(obj interface{}, fh io.Reader) error {
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(blob, obj); err != nil {
		return err
	}
	return nil
}

func (c *Manager) Write(obj interface{}, fh io.Writer) error {
	blob, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	fh.Write(blob)
	fh.Write([]byte("\n"))
	return nil
}
