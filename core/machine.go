// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"

	"github.com/munbot/master/config"
)

type Machine interface {
	SetState(StateID) error
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
	Config() *config.Config
	ConfigFlags() *config.Flags
	CoreFlags() *Flags
}
