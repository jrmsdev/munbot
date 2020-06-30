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
}

type baseValue struct {
	t string
	n string
}

func newValue(t, n string) *baseValue {
	return &baseValue{t, n}
}

func (v *baseValue) Name() string {
	return v.n
}

func (v *baseValue) Type() string {
	return v.t
}
