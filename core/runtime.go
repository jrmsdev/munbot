// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
)

type Runtime interface {
	Init(context.Context) (context.Context, error)
	Configure(*flags.Flags, *config.Flags, *config.Config) error
	Start() error
	Stop() error
}
