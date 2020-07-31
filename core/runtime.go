// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"

	"github.com/munbot/master/config"
)

type Runtime interface {
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
	Init() error
	Configure(*Flags, *config.Flags, *config.Config) error
	Start() error
	Stop() error
}
