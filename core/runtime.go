// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"

	"github.com/munbot/master/config"
)

type Runtime interface {
	String() string
	UUID() string
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
	Lock(context.Context) (context.Context, error)
	Configure(*config.Config, *config.Flags, *Flags) error
}
