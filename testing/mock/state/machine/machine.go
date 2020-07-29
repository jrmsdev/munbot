// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package machine

import (
	"context"
	"errors"
	"flag"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/state"
)

var _ state.Machine = &MockSM{}

// MockSM wraps state.SM so we can mock some parts of it.
// You can pass nil or empty values to most of the method calls and so default
// ones will be created. Like in example, you can pass nil as a ctx arg for a
// method an a context.Background() will be created be to call the real method.
type MockSM struct {
	*state.SM
	WithInitError bool
	WithRunError  bool
}

// NewMockSM creates a new mockable state.Machine implementation.
func NewMockSM() *MockSM {
	return &MockSM{SM: state.NewMachine().(*state.SM)}
}

// Init calls real method or returns an error if WithInitError.
// If config.Flags or core.Flags args are nil, New() ones will be created.
func (m *MockSM) Init(cf *config.Flags, fl *core.Flags) error {
	if m.WithInitError {
		return errors.New("mock init error")
	}
	if cf == nil {
		fs := flag.NewFlagSet("mock-state-machine", flag.PanicOnError)
		cf = config.NewFlags(fs)
	}
	if fl == nil {
		fl = core.NewFlags()
	}
	return m.SM.Init(cf, fl)
}

// Run calls real method or returns an error if WithRunError.
// If ctx is nil, context.Background() will be used.
func (m *MockSM) Run(ctx context.Context) error {
	if m.WithRunError {
		return errors.New("mock run error")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return m.SM.Run(ctx)
}