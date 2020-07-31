// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
	"errors"
)

const lockKey ctxkey = 0

var ErrLock error = errors.New("core: context lock failed")

func (k *Core) lock(ctx context.Context) (context.Context, error) {
	if err := k.tryLock(); err != nil {
		return ctx, err
	}
	k.ctx = context.WithValue(ctx, lockKey, k.locked)
	return k.ctx, nil
}

func (k *Core) tryLock() error {
	if k.mu.TryLock(nil) {
		k.locked = k.uuid
		return nil
	}
	return ErrLock
}
