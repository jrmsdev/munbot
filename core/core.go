// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"
	"time"

	"github.com/munbot/master/config"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

type key int

const lockKey key = 0

func lockContext(rt *Core) (context.Context, error) {
	var err error
	rt.uuid, err = tryLock(rt.mu)
	if err != nil {
		return nil, err
	}
	return context.WithValue(rt.ctx, lockKey, rt.uuid), nil
}

func tryLock(mu *lock.Locker) (string, error) {
	if mu.TryLockTimeout(time.Second) {
		return uuid.Rand(), nil
	}
	return "", errors.New("core lock timeout")
}

type Core struct {
	ctx      context.Context
	mu       *lock.Locker
	uuid     string
	cfg      *config.Config
	cfgFlags *config.Flags
	flags    *Flags
}

func NewRuntime(ctx context.Context) Runtime {
	return &Core{ctx: ctx, mu: lock.New()}
}

func (rt *Core) String() string {
	return "core.runtime:" + rt.uuid
}

func (rt *Core) UUID() string {
	return rt.uuid
}

func (rt *Core) Context() context.Context {
	return rt.ctx
}

func (rt *Core) Lock() error {
	var err error
	rt.ctx, err = lockContext(rt)
	return err
}

func (rt *Core) Configure(cfg *config.Config, cfl *config.Flags, f *Flags) error {
	if rt.uuid == "" {
		return errors.New("core runtime not locked")
	}
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	default:
		rt.cfg = cfg
		rt.cfgFlags = cfl
		rt.flags = f
	}
	return nil
}
