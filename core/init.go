// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"context"
)

func (rt *Core) Init(ctx context.Context) (context.Context, error) {
	select {
	case <-ctx.Done():
		return ctx, ctx.Err()
	}
	var err error
	ctx, err = rt.WithContext(ctx)
	if err != nil {
		return ctx, err
	}
	err = rt.state.Init()
	return ctx, err
}
