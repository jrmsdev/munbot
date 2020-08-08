// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"

	"github.com/munbot/master/api"
	"github.com/munbot/master/console"
	"github.com/munbot/master/robot/master"
	"github.com/munbot/master/utils/lock"
)

var ErrMemLock error = errors.New("core: mem lock failed")

type Mem struct {
	mu      *lock.Locker
	Master  master.Munbot
	Api     api.Server
	Console console.Server
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
