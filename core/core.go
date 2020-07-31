// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

var _ Runtime = &Core{}

type Core struct {
	ctx      context.Context
	mu       *lock.Locker
	uuid     string
	locked   string
	cfg      *config.Config
	cfgFlags *config.Flags
	flags    *Flags
	state    State
}

func NewRuntime() Runtime {
	return &Core{mu: lock.New(), uuid: uuid.Rand()}
}

func (rt *Core) String() string {
	return "Core:" + rt.uuid
}

func (rt *Core) UUID() string {
	return rt.uuid
}

func (rt *Core) Init(ctx context.Context) (context.Context, error) {
	select {
	case <-ctx.Done():
		return ctx, ctx.Err()
	}
	var err error
	ctx, err = rt.WithContext(ctx)
	if err != nil {
		return ctx, err
	}
	err = rt.state.Init()
	return ctx, err
}

var ErrCtxNoLock error = errors.New("core: no context locked")

func (rt *Core) Configure(kfl *Flags, cfl *config.Flags, cfg *config.Config) error {
	if rt.locked == "" || rt.ctx == nil {
		return ErrCtxNoLock
	}
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	default:
		rt.cfg = cfg
		rt.cfgFlags = cfl
		rt.flags = kfl
	}
	// TODO: read config, parse flags, etc...
	return rt.state.Configure(rt.flags, rt.cfgFlags, rt.cfg)
}

func (rt *Core) Start() error {
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	}
	return rt.state.Start()
}

func (rt *Core) Stop() error {
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	}
	return rt.state.Stop()
}
