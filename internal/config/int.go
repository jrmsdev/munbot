// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/log"
	"strconv"
)

type IntValue struct {
	*baseValue
	i int
}

func (v *IntValue) String() string {
	return strconv.Itoa(v.i)
}

func (v *IntValue) Value() int {
	return v.i
}

func (v *IntValue) Update(newval string) error {
	log.Debugf("update %s:%s", v.Type(), v.Name())
	var err error
	v.i, err = strconv.Atoi(newval)
	v.setDirty()
	return err
}

func (v *IntValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	return v.Update(string(b))
}

func (v *IntValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return []byte(v.String()), nil
}