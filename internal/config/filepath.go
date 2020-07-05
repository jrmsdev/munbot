// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type AbsFilepath struct {
	Value string `json:"-"`
}

func (p *AbsFilepath) String() string {
	return p.Value
}

func (p *AbsFilepath) UnmarshalJSON(b []byte) error {
	return unmarshalFilepath(&p.Value, b, absPath)
}

func (p *AbsFilepath) MarshalJSON() ([]byte, error) {
	return marshalFilepath(p.Value, absPath)
}

type RelFilepath struct {
	Value string `json:"-"`
}

func (p *RelFilepath) String() string {
	return p.Value
}

func (p *RelFilepath) UnmarshalJSON(b []byte) error {
	return unmarshalFilepath(&p.Value, b, relPath)
}

func (p *RelFilepath) MarshalJSON() ([]byte, error) {
	return marshalFilepath(p.Value, relPath)
}

func unmarshalFilepath(p interface{}, b []byte, pt pathType) error {
	var s string
	if err := json.Unmarshal(b, s); err != nil {
		return fmt.Errorf("filepath: %s", err)
	}
	p = filepath.Clean(filepath.FromSlash(s))
	return checkFilepath(pt, p.(string))
}

func marshalFilepath(p interface{}, pt pathType) ([]byte, error) {
	fp := filepath.Clean(filepath.ToSlash(p.(string)))
	if err := checkFilepath(pt, fp); err != nil {
		return nil, err
	}
	blob, err := json.Marshal(&fp)
	if err != nil {
		return nil, fmt.Errorf("filepath: %s", err)
	}
	return blob, nil
}

func checkFilepath(pt pathType, p string) error {
	if pt == absPath {
		return checkAbsFilepath(p)
	}
	return checkRelFilepath(p)
}

func checkAbsFilepath(p string) error {
	if !filepath.IsAbs(p) {
		return fmt.Errorf("%s: should be an absolute filepath", p)
	}
	return nil
}

func checkRelFilepath(p string) error {
	if filepath.IsAbs(p) {
		return fmt.Errorf("%s: should be a relative filepath", p)
	}
	return nil
}
