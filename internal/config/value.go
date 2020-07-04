// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"

	"github.com/jrmsdev/munbot/log"
)

type Value interface {
	json.Marshaler
	json.Unmarshaler
	Name() string
	String() string
	Type() string
	Update(string) error
	modified() bool
}

type baseValue struct {
	t     string
	n     string
	dirty bool
}

func newValue(t, n string) *baseValue {
	return &baseValue{t, n, false}
}

func (v *baseValue) Name() string {
	return v.n
}

func (v *baseValue) Type() string {
	return v.t
}

func (v *baseValue) setDirty(old, new string) bool {
	if new == "null" {
		return false
	}
	if new != old {
		v.dirty = true
	}
	return v.dirty
}

func (v *baseValue) modified() bool {
	return v.dirty
}

func (v *baseValue) jsonUnmarshal(blob []byte, dst interface{}) error {
	log.Debugf("json unmarshal %s %s", v.t, v.n)
	return json.Unmarshal(blob, dst)
}

func (v *baseValue) jsonMarshal(src interface{}) ([]byte, error) {
	log.Debugf("json marshal %s %s", v.t, v.n)
	return json.Marshal(src)
}
