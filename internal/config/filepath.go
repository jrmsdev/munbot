// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"
	"strconv"

	"github.com/jrmsdev/munbot/log"
)

type FilepathValue struct {
	*baseValue
	p string
}

func (v *FilepathValue) String() string {
	return v.p
}

func (v *FilepathValue) Value() string {
	return v.p
}

func (v *FilepathValue) Update(newval string) error {
	log.Debugf("update %s:%s", v.Type(), v.Name())
	v.setDirty()
	v.p = filepath.Clean(filepath.FromSlash(newval))
	return nil
}

func (v *FilepathValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return v.Update(s)
}

func (v *FilepathValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	fp := filepath.Clean(filepath.ToSlash(v.p))
	return []byte(strconv.Quote(fp)), nil
}
