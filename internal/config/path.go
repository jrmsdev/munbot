// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"path"
)

type pathType int
const absPath pathType = 1
const relPath pathType = 2

type AbsPath struct {
	Value string `json:"-"`
}

func (p *AbsPath) String() string {
	return p.Value
}

func (p *AbsPath) UnmarshalJSON(b []byte) error {
	return unmarshalPath(&p.Value, b, absPath)
}

func (p *AbsPath) MarshalJSON() ([]byte, error) {
	return marshalPath(p.Value, absPath)
}

type RelPath struct {
	Value string `json:"-"`
}

func (p *RelPath) String() string {
	return p.Value
}

func (p *RelPath) UnmarshalJSON(b []byte) error {
	return unmarshalPath(&p.Value, b, relPath)
}

func (p *RelPath) MarshalJSON() ([]byte, error) {
	return marshalPath(p.Value, relPath)
}

func unmarshalPath(p interface{}, b []byte, pt pathType) error {
	var s string
	if err := json.Unmarshal(b, s); err != nil {
		return fmt.Errorf("path: %s", err)
	}
	p = path.Clean(s)
	return checkPath(pt, p.(string))
}

func marshalPath(p interface{}, pt pathType) ([]byte, error) {
	s := path.Clean(p.(string))
	if err := checkPath(pt, s); err != nil {
		return nil, err
	}
	blob, err := json.Marshal(&s)
	if err != nil {
		return nil, fmt.Errorf("path: %s", err)
	}
	return blob, nil
}

func checkPath(pt pathType, p string) error {
	if pt == absPath {
		return checkAbsPath(p)
	}
	return checkRelPath(p)
}

func checkAbsPath(p string) error {
	if !path.IsAbs(p) {
		return fmt.Errorf("%s: should be an absolute path", p)
	}
	return nil
}

func checkRelPath(p string) error {
	if path.IsAbs(p) {
		return fmt.Errorf("%s: should be a relative path", p)
	}
	return nil
}
