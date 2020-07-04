// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
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
	if v.setDirty(v.String(), newval) {
		var err error
		v.b, err = strconv.ParseBool(newval)
		return err
	}
	return nil
}

func (v *BoolValue) UnmarshalJSON(b []byte) error {
	var d bool
	if err := v.jsonUnmarshal(b, &d); err != nil {
		return err
	}
	return v.Update(string(b))
}

func (v *BoolValue) MarshalJSON() ([]byte, error) {
	return v.jsonMarshal(&v.b)
}

func (v *BoolValue) IsTrue() bool {
	return v.b == true
}

func (v *BoolValue) IsFalse() bool {
	return v.b == false
}
