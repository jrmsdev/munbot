// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"

	"github.com/munbot/master/internal/api"
	"github.com/munbot/master/internal/auth"
	"github.com/munbot/master/internal/console"
	"github.com/munbot/master/robot/master"
	"github.com/munbot/master/utils/lock"
)

var ErrMemLock error = errors.New("core: mem lock failed")

type Mem struct {
	mu      *lock.Locker
	Auth    auth.Manager
	Api     api.Server
	Console console.Server
	Master  master.Munbot
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
