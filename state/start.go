// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
)

var _ State = &StartState{}

type StartState struct {
	m   *SM
	id  StateID
	err error
}

func newStart(m *SM) *StartState {
	return &StartState{m: m, id: Start}
}

func (s *StartState) String() string {
	return s.id.String()
}

func (s *StartState) Error() error {
	return s.err
}

func (s *StartState) Run(ctx context.Context) (context.Context, Status) {
	select {
	case <-ctx.Done():
		return ctx, DONE
	default:
	}
	return ctx, EXIT
}
