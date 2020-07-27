// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"
)

var ConfigureError = errors.New("internal error: run Master.Configure first")

type ConfigureState struct {
	m   *Machine
	err error
}

func newConfigure(m *Machine) *ConfigureState {
	return &ConfigureState{m: m}
}

func (s *ConfigureState) String() string {
	return "Configure"
}

func (s *ConfigureState) Error() error {
	return s.err
}

func (s *ConfigureState) Run(ctx context.Context) Status {
	select {
	case <-ctx.Done():
		return DONE
	default:
	}
	return EXIT
}
