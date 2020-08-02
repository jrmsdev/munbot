// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
)

type ctxkey int

func (k *Core) Context() context.Context {
	return k.ctx
}

func (k *Core) WithContext(ctx context.Context) (context.Context, error) {
	var err error
	lockId := k.fromContext(ctx, lockKey)
	if lockId == "" {
		if k.locked != "" {
			// we are locked but the new context is not, so copy our lock key to
			// the new context
			ctx = k.ctxCopy(ctx, lockKey)
		} else {
			// neither we nor the context is locked... so go for it
			if ctx, err = k.lock(ctx); err != nil {
				return ctx, err
			}
		}
	} else {
		if k.locked == "" {
			// the new context is locked but we are not, try to lock ourselves
			if ctx, err = k.lock(ctx); err != nil {
				return ctx, err
			}
		}
		// use new context lock key for future validations
		k.locked = lockId
	}
	k.ctx = ctx
	return k.ctx, nil
}

func (k *Core) ctxCopy(dst context.Context, key ctxkey) context.Context {
	if v, ok := k.ctx.Value(key).(string); ok {
		dst = context.WithValue(dst, key, v)
	}
	return dst
}

func (k *Core) fromContext(ctx context.Context, key ctxkey) string {
	s, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return s
}
