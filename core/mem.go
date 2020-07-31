// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/utils/lock"
)

var mem *Mem

func init() {
	mem = newMem()
}

var ErrMemLock error = errors.New("core: mem lock failed")

type Mem struct {
	mu       *lock.Locker
	Flags    *Flags
	Cfg      *config.Config
	CfgFlags *config.Flags
}

func newMem() *Mem {
	return &Mem{mu: lock.New()}
}

func (m *Mem) Lock() error {
	if m.mu.TryLock(nil) {
		return nil
	}
	return ErrMemLock
}

func (m *Mem) Unlock() {
	m.mu.Unlock()
}
