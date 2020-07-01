// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/log"
	"strconv"
)

type BoolValue struct {
	*baseValue
	b bool
}

func (v *BoolValue) String() string {
	return strconv.FormatBool(v.b)
}

func (v *BoolValue) Value() bool {
	return v.b
}

func (v *BoolValue) Update(newval string) error {
	log.Debugf("update %s:%s", v.Type(), v.Name())
	var err error
	v.b, err = strconv.ParseBool(newval)
	v.setDirty()
	return err
}

func (v *BoolValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	return v.Update(string(b))
}

func (v *BoolValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return []byte(v.String()), nil
}

func (v *BoolValue) IsTrue() bool {
	return v.b == true
}

func (v *BoolValue) IsFalse() bool {
	return v.b == false
}
