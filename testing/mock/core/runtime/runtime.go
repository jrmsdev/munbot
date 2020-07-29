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
// If MockContext is set, it will be "loaded" before calling any real method.
// But it will be loaded only once, then MockContext is set as nil.
type MockRuntime struct {
	*core.Core
	ctx                context.Context
	MockContext        context.Context
	WithLockError      bool
	WithConfigureError bool
}

// NewMockRuntime creates a new mockable runtime with a Background context.
func NewMockRuntime() *MockRuntime {
	ctx := context.Background()
	return &MockRuntime{
		Core: core.NewRuntime(ctx).(*core.Core),
		ctx:  ctx,
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
func (rt *MockRuntime) Lock() error {
	if rt.WithLockError {
		return errors.New("mock lock error")
	}
	return rt.Core.Lock()
}

// Configure calls real Configure method or returns an error if WithConfigureError.
func (rt *MockRuntime) Configure(cfg *config.Config, cfl *config.Flags, f *core.Flags) error {
	if rt.WithConfigureError {
		return errors.New("mock configure error")
	}
	rt.withContext()
	return rt.Core.Configure(cfg, cfl, f)
}
