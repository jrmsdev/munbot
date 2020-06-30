// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"errors"
	"fmt"
)

type registry map[string]Value

func (r registry) Dump() {
	for k, v := range r {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func (r registry) Update(key, newval string) error {
	if _, ok := r[key]; !ok {
		return errors.New(fmt.Sprintf("invalid config key: %s", key))
	}
	return r[key].Update(newval)
}
