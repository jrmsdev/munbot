// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"
	"github.com/jrmsdev/munbot/log"
)

type IntValue struct {
	*baseValue
	i int
}

func (v *IntValue) String() string {
	return strconv.Itoa(v.i)
}

func (v *IntValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	var err error
	v.i, err = strconv.Atoi(string(b))
	return err
}

func (v *IntValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	return []byte(v.String()), nil
}

func (v *IntValue) Value() int {
	return v.i
}
