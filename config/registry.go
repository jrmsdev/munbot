// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"fmt"
)

type registry map[string]Value

func (r *registry) Dump() {
	for k, v := range *r {
		fmt.Printf("%s=%s\n", k, v)
	}
}
