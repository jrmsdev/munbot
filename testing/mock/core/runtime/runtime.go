// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package runtime

import (
	"context"
	"errors"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
)

var _ core.Runtime = &MockRuntime{}

// MockRuntime runs a real core runtime but you can mock some parts of it.
// If MockContext is set, it will be "loaded" before calling any real method
// that does not receive a ctx arg, but it uses the "attached" context.
// But it will be loaded only once, then MockContext is set as nil.
type MockRuntime struct {
	*core.Core
	MockContext        context.Context
	WithLockError      bool
	WithConfigureError bool
}

// NewMockRuntime creates a new mockable runtime with a Background context.
func NewMockRuntime() *MockRuntime {
	return &MockRuntime{
		Core: core.NewRuntime().(*core.Core),
	}
}

func (rt *MockRuntime) withContext() {
	if rt.MockContext != nil {
		if _, err := rt.Core.WithContext(rt.MockContext); err != nil {
			panic(err)
		}
		// just once
		rt.MockContext = nil
	}
}

// Lock calls real Lock method or returns an error if WithLockError is true.
func (rt *MockRuntime) Lock(ctx context.Context) (context.Context, error) {
	if rt.WithLockError {
		return ctx, errors.New("mock lock error")
	}
	return rt.Core.Lock(ctx)
}

// Configure calls real Configure method or returns an error if WithConfigureError.
func (rt *MockRuntime) Configure(cfg *config.Config, cfl *config.Flags, f *core.Flags) error {
	if rt.WithConfigureError {
		return errors.New("mock configure error")
	}
	rt.withContext()
	return rt.Core.Configure(cfg, cfl, f)
}
