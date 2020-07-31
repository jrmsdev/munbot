// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"

	"github.com/munbot/master/config"
)

type Machine interface {
	SetState(ID)
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
	Config() *config.Config
	ConfigFlags() *config.Flags
}
