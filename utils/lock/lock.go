// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package lock wraps github.com/subchen/go-trylock package.
package lock

import (
	"github.com/subchen/go-trylock/v2"
)

// Locker wraps trylock.TryLocker.
type Locker struct {
	trylock.TryLocker
}

// New creates a new trylock.TryLocker.
func New() *Locker {
	return &Locker{trylock.New()}
}
