// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
)

type Machine interface {
	SetState(StateID)
	Context() context.Context
	WithContext(context.Context) (context.Context, error)
}
