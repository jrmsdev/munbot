// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
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

func (v *baseValue) setDirty() {
	v.dirty = true
}

func (v *baseValue) modified() bool {
	return v.dirty
}
