// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
)

type StartState struct {
	m   *Machine
	err error
}

func newStart(m *Machine) *StartState {
	return &StartState{m: m}
}

func (s *StartState) String() string {
	return "Start"
}

func (s *StartState) Error() error {
	return s.err
}

func (s *StartState) Run(ctx context.Context) Status {
	select {
	case <-ctx.Done():
		return DONE
	default:
	}
	return EXIT
}
