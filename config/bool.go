// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"
	"github.com/jrmsdev/munbot/log"
)

type BoolValue struct {
	*baseValue
	b bool
}

func NewBool(name string, defval bool) *BoolValue {
	return &BoolValue{newValue("bool", name), defval}
}

func (v *BoolValue) String() string {
	return strconv.FormatBool(v.b)
}

func (v *BoolValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	var err error
	v.b, err = strconv.ParseBool(string(b))
	return err
}

func (v *BoolValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return []byte(v.String()), nil
}

func (v *BoolValue) Value() bool {
	return v.b
}
