// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"

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
	log.Debugf("update %s:%s", v.Type(), v.Name())
	v.s = newval
	v.setDirty()
	return nil
}

func (v *StringValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return v.Update(s)
}

func (v *StringValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return []byte(strconv.Quote(v.s)), nil
}
