// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

var _ Runtime = &Core{}
var _ Machine = &Core{}

type Core struct {
	rt       *Mem
	ctx      context.Context
	mu       *lock.Locker
	uuid     string
	locked   string
	state    State
	stid     StateID
	sInit    State
}

func NewRuntime() Runtime {
	k := &Core{
		rt: newMem(),
		mu: lock.New(),
		uuid: uuid.Rand(),
	}
	k.sInit = NewInit(k, k.rt)
	k.SetState(Init)
	return k
}

func (k *Core) String() string {
	return "Core:" + k.uuid
}

func (k *Core) UUID() string {
	return k.uuid
}

func (k *Core) SetState(s StateID) {
	log.Debugf("state %s set %s", StateName(k.stid), StateName(s))
	if s == k.stid {
		log.Panicf("core: state %s set twice", StateName(s))
	}
	switch s {
	case Init:
		k.state = k.sInit
	default:
		log.Panicf("core: set %s", StateName(s))
	}
	k.stid = s
}

func (k *Core) Init(ctx context.Context) (context.Context, error) {
	select {
	case <-ctx.Done():
		return ctx, ctx.Err()
	}
	if err := k.state.Init(); err != nil {
		return ctx, err
	}
	return k.WithContext(ctx)
}

var ErrCtxNoLock error = errors.New("core: no context locked")

func (k *Core) Configure(kfl *Flags, cfl *config.Flags, cfg *config.Config) error {
	if k.locked == "" || k.ctx == nil {
		return ErrCtxNoLock
	}
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	default:
		// TODO: read config, parse flags, etc...
		if err := k.state.Configure(); err != nil {
			return err
		}
	}
	k.rt.Flags = kfl
	k.rt.CfgFlags = cfl
	k.rt.Cfg = cfg
	return nil
}

func (k *Core) Start() error {
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	}
	return k.state.Start()
}

func (k *Core) Stop() error {
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	}
	return k.state.Stop()
}
