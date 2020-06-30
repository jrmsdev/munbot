// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"errors"
	"fmt"
)

type registry struct {
	db map[string]Value
}

func newReg() *registry {
	return &registry{make(map[string]Value)}
}

func (r *registry) Dump() {
	for k, v := range r.db {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func (r *registry) Update(key, newval string) error {
	if _, ok := r.db[key]; !ok {
		return errors.New(fmt.Sprintf("invalid config key: %s", key))
	}
	return r.db[key].Update(newval)
}
