// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"fmt"
	"path"
	"strconv"

	"github.com/jrmsdev/munbot/log"
)

type PathValue struct {
	*baseValue
	p string
	forceAbs bool
}

func (v *PathValue) String() string {
	return v.p
}

func (v *PathValue) Value() string {
	return v.p
}

func (v *PathValue) Update(newval string) error {
	log.Debugf("update %s:%s", v.Type(), v.Name())
	if v.setDirty(v.p, newval) {
		v.p = path.Clean(newval)
		return v.check(v.p)
	}
	return nil
}

func (v *PathValue) UnmarshalJSON(b []byte) error {
	log.Debugf("json unmarshal %s:%s", v.Type(), v.Name())
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return v.Update(s)
}

func (v *PathValue) MarshalJSON() ([]byte, error) {
	log.Debugf("json marshal %s:%s", v.Type(), v.Name())
	p := path.Clean(v.p)
	if err := v.check(p); err != nil {
		return nil, err
	}
	return []byte(strconv.Quote(p)), nil
}

func (v *PathValue) check(p string) error {
	if v.forceAbs {
		return v.checkAbs(p)
	}
	return v.checkRel(p)
}

func (v *PathValue) checkAbs(p string) error {
	if !path.IsAbs(p) {
		return fmt.Errorf("%s option should be an absolute path: %s",
			v.Name(), p)
	}
	return nil
}

func (v *PathValue) checkRel(p string) error {
	if path.IsAbs(p) {
		return fmt.Errorf("%s option should be a relative path: %s",
			v.Name(), p)
	}
	return nil
}
