// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/state"
	"github.com/munbot/master/log"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

var _ Runtime = &Core{}
var _ state.Machine = &Core{}

type Core struct {
	ctx      context.Context
	mu       *lock.Locker
	uuid     string
	locked   string
	cfg      *config.Config
	cfgFlags *config.Flags
	flags    *Flags
	state    state.State
	stid     state.ID
	sInit    state.State
}

func NewRuntime() Runtime {
	rt := &Core{
		mu: lock.New(),
		uuid: uuid.Rand(),
	}
	rt.sInit = state.NewInit(rt)
	rt.SetState(state.Init)
	return rt
}

func (rt *Core) String() string {
	return "Core:" + rt.uuid
}

func (rt *Core) UUID() string {
	return rt.uuid
}

func (rt *Core) Config() *config.Config {
	return rt.cfg
}

func (rt *Core) ConfigFlags() *config.Flags {
	return rt.cfgFlags
}

func (rt *Core) SetState(s state.ID) {
	log.Debugf("state %s set %s", state.Name(rt.stid), state.Name(s))
	if s == rt.stid {
		log.Panicf("core: state %s set twice", state.Name(s))
	}
	switch s {
	case state.Init:
		rt.state = rt.sInit
	default:
		log.Panicf("core: set %s", state.Name(s))
	}
	rt.stid = s
}

func (rt *Core) Init(ctx context.Context) (context.Context, error) {
	select {
	case <-ctx.Done():
		return ctx, ctx.Err()
	}
	if err := rt.state.Init(); err != nil {
		return ctx, err
	}
	return rt.WithContext(ctx)
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
		// TODO: read config, parse flags, etc...
		if err := rt.state.Configure(); err != nil {
			return err
		}
	}
	rt.flags = kfl
	rt.cfgFlags = cfl
	rt.cfg = cfg
	return nil
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
