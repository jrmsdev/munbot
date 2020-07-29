// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
)

var _ State = &StartState{}

type StartState struct {
	m   *SM
	err error
}

func newStart(m *SM) *StartState {
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
