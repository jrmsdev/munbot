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

func (v *StringValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	v.s = string(b)
	return nil
}

func (v *StringValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	if v.s == "" {
		v.s = `""`
	}
	return []byte(v.s), nil
}

func (v *StringValue) Value() string {
	return v.s
}
