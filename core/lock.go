// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
	"errors"
	"time"
)

const lockKey ctxkey = 0

var ErrLock error = errors.New("core: context lock failed")

func (rt *Core) lock(ctx context.Context) (context.Context, error) {
	if err := rt.tryLock(); err != nil {
		return ctx, err
	}
	rt.ctx = context.WithValue(ctx, lockKey, rt.locked)
	return rt.ctx, nil
}

func (rt *Core) tryLock() error {
	if rt.mu.TryLockTimeout(300 * time.Millisecond) {
		rt.locked = rt.uuid
		return nil
	}
	return ErrLock
}
