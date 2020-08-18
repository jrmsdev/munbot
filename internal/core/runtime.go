// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

type Runtime interface {
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
	Init(context.Context, *config.Flags, *config.Config) (context.Context, error)
	Configure() error
	Start() error
	Stop() error
}

func (k *Core) Context() context.Context {
	return k.ctx
}

func (k *Core) WithContext(ctx context.Context) (context.Context, error) {
	if err := k.rt.Lock(); err != nil {
		return ctx, err
	}
	defer k.rt.Unlock()
	k.ctx = ctx
	return k.ctx, nil
}

func (k *Core) Init(ctx context.Context, cfl *config.Flags, cfg *config.Config) (context.Context, error) {
	log.Debugf("[%s] Init", k.State())
	var err error
	ctx, err = k.WithContext(ctx)
	if err != nil {
		return ctx, k.error(err)
	}
	select {
	case <-ctx.Done():
		return ctx, k.error(ctx.Err())
	default:
		if err := k.rt.Lock(); err != nil {
			return ctx, k.error(err)
		}
	}
	defer k.rt.Unlock()
	k.cfg = cfg
	k.cfl = cfl
	if err = k.state.Init(); err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (k *Core) Configure() error {
	log.Debugf("[%s] Configure", k.State())
	select {
	case <-k.ctx.Done():
		return k.error(k.ctx.Err())
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Configure(); err != nil {
		return err
	}
	return nil
}

func (k *Core) Start() error {
	log.Debugf("[%s] Start", k.State())
	select {
	case <-k.ctx.Done():
		return k.error(k.ctx.Err())
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Start(); err != nil {
		return err
	}
	if err := k.state.Run(); err != nil {
		return err
	}
	log.Debug("bye!")
	return nil
}

func (k *Core) Stop() error {
	log.Debugf("[%s] Stop", k.State())
	select {
	case <-k.ctx.Done():
		return k.error(k.ctx.Err())
	default:
		if err := k.rt.Lock(); err != nil {
			return k.error(err)
		}
	}
	defer k.rt.Unlock()
	if err := k.state.Stop(); err != nil {
		return err
	}
	if err := k.state.Halt(); err != nil {
		return err
	}
	return nil
}
