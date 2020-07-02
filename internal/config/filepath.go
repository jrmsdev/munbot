// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/jrmsdev/munbot/log"
)

type FilepathValue struct {
	*baseValue
	p string
	forceAbs bool
}

func (v *FilepathValue) String() string {
	return v.p
}

func (v *FilepathValue) Value() string {
	return v.p
}

func (v *FilepathValue) Update(newval string) error {
	log.Debugf("update %s:%s", v.Type(), v.Name())
	fp := filepath.Clean(filepath.FromSlash(newval))
	if v.setDirty(v.p, fp) {
		v.p = fp
		return v.check(v.p)
	}
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
	if err := v.check(fp); err != nil {
		return nil, err
	}
	return []byte(strconv.Quote(fp)), nil
}

func (v *FilepathValue) check(p string) error {
	if v.forceAbs {
		return v.checkAbs(p)
	}
	return v.checkRel(p)
}

func (v *FilepathValue) checkAbs(p string) error {
	if !filepath.IsAbs(p) {
		return fmt.Errorf("%s option should be an absolute filepath: %s",
			v.Name(), p)
	}
	return nil
}

func (v *FilepathValue) checkRel(p string) error {
	if filepath.IsAbs(p) {
		return fmt.Errorf("%s option should be a relative filepath: %s",
			v.Name(), p)
	}
	return nil
}
