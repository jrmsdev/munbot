// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

import (
	"context"

	"github.com/munbot/master/core/lock"
)

var rt *runtime

func init() {
	rt = newRuntime()
}

type runtime struct {
}

func newRuntime() *runtime {
	return &runtime{}
}

func Lock(ctx context.Context) (context.Context, error) {
	return lock.NewContext(ctx)
}
