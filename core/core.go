// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

var _ Runtime = &Core{}
var _ Machine = &Core{}

type Core struct {
	rt         *Mem
	ctx        context.Context
	mu         *lock.Locker
	uuid       string
	locked     string
	state      State
	stid       StateID
	sInit      State
	sConfigure State
	sRun       State
}

func NewRuntime() Runtime {
	return New(mem)
}

func New(m *Mem) *Core {
	k := &Core{
		rt:   m,
		mu:   lock.New(),
		uuid: uuid.Rand(),
	}
	k.sInit = newInit(k, k.rt)
	k.sConfigure = newConfigure(k, k.rt)
	k.sRun = newRun(k, k.rt)
	k.SetState(Init)
	return k
}

func (k *Core) String() string {
	return "Core:" + k.uuid
}

func (k *Core) UUID() string {
	return k.uuid
}

func (k *Core) error(err error) error {
	log.Output(2, fmt.Sprintf("[ERROR] core %s: %s", StateName(k.stid), err))
	return err
}

func (k *Core) errorf(f string, args ...interface{}) error {
	err := errors.New(fmt.Sprintf(f, args...))
	log.Output(2, fmt.Sprintf("[ERROR] core %s: %s", StateName(k.stid), err))
	return err
}

func (k *Core) SetState(s StateID) error {
	log.Debugf("state %s set %s", StateName(k.stid), StateName(s))
	if s == k.stid {
		return k.errorf("core: state %s set twice", StateName(s))
	}
	switch s {
	case Init:
		k.state = k.sInit
	case Configure:
		k.state = k.sConfigure
	case Run:
		k.state = k.sRun
	default:
		k.errorf("core: set %s", StateName(s))
	}
	k.stid = s
	return nil
}

func (k *Core) Init(ctx context.Context) (context.Context, error) {
	log.Debug("init")
	select {
	case <-ctx.Done():
		return ctx, ctx.Err()
	default:
		if err := k.rt.Lock(); err != nil {
			return ctx, k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Init(); err != nil {
		return ctx, k.error(err)
	}
	var err error
	ctx, err = k.WithContext(ctx)
	if err != nil {
		return ctx, k.error(err)
	}
	return ctx, nil
}

var ErrCtxNoLock error = errors.New("core: no context locked")

func (k *Core) Configure(kfl *Flags, cfl *config.Flags, cfg *config.Config) error {
	log.Debug("configure")
	if k.locked == "" || k.ctx == nil {
		return k.error(ErrCtxNoLock)
	}
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Configure(); err != nil {
		return k.error(err)
	}
	return nil
}

func (k *Core) Start() error {
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Start(); err != nil {
		return k.error(err)
	}
	return nil
}

func (k *Core) Stop() error {
	select {
	case <-k.ctx.Done():
		return k.ctx.Err()
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Stop(); err != nil {
		return k.error(err)
	}
	return nil
}
