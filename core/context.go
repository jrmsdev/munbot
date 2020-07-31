// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
)

type ctxkey int

func (rt *Core) Context() context.Context {
	return rt.ctx
}

func (rt *Core) WithContext(ctx context.Context) (context.Context, error) {
	var err error
	lockId := rt.fromContext(ctx, lockKey)
	if lockId == "" {
		if rt.locked != "" {
			// we are locked but the new context is not, so copy our lock key to
			// the new context
			ctx = rt.ctxCopy(ctx, lockKey)
		} else {
			// neither we nor the context is locked... so go for it
			if ctx, err = rt.lock(ctx); err != nil {
				return ctx, err
			}
		}
	} else {
		if rt.locked == "" {
			// the new context is locked but we are not, try to lock ourselves
			if ctx, err = rt.lock(ctx); err != nil {
				return ctx, err
			}
		}
		// use new context lock key for future validations
		rt.locked = lockId
	}
	rt.ctx = ctx
	return rt.ctx, nil
}

func (rt *Core) ctxCopy(dst context.Context, k ctxkey) context.Context {
	if v, ok := rt.ctx.Value(k).(string); ok {
		dst = context.WithValue(dst, k, v)
	}
	return dst
}

func (rt *Core) fromContext(ctx context.Context, k ctxkey) string {
	s, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return s
}
