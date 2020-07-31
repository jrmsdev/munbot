// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"

	"github.com/munbot/master/config"
)

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
