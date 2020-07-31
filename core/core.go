// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"

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
