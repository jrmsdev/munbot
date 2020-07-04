// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
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
	if v.setDirty(v.String(), newval) {
		var err error
		v.i, err = strconv.Atoi(newval)
		return err
	}
	return nil
}

func (v *IntValue) UnmarshalJSON(b []byte) error {
	var i int
	if err := v.jsonUnmarshal(b, &i); err != nil {
		return err
	}
	return v.Update(strconv.Itoa(i))
}

func (v *IntValue) MarshalJSON() ([]byte, error) {
	return v.jsonMarshal(&v.i)
}
