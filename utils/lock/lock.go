// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Lock utils.
package lock

import (
	"github.com/subchen/go-trylock/v2"
)

type Locker struct {
	trylock.TryLocker
}

func New() *Locker {
	return &Locker{trylock.New()}
}
