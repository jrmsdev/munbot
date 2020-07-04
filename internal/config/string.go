// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/log"
)

type StringValue struct {
	*baseValue
	s string
}

func (v *StringValue) String() string {
	return v.s
}

func (v *StringValue) Value() string {
	return v.s
}

func (v *StringValue) Update(newval string) error {
	if v.setDirty(v.s, newval) {
		v.s = newval
	}
	return nil
}

func (v *StringValue) UnmarshalJSON(b []byte) error {
	var s string
	if err := v.jsonUnmarshal(b, &s); err != nil {
		return err
	}
	return v.Update(s)
}

func (v *StringValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return v.jsonMarshal(&v.s)
}
