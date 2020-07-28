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
	Lock() error
	Configure(*config.Config, *config.Flags, *Flags) error
}
