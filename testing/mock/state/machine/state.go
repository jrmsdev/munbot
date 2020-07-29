// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package machine

import (
	"context"
	"errors"

	"github.com/munbot/master/state"
)

var _ state.State = &MockState{}

// FIXME: MockState not implemented/used.
type MockState struct {
	ExitStatus state.Status
}

func newMockState() *MockState {
	return &MockState{ExitStatus: state.OK}
}

func (s *MockState) Error() error {
	if s.ExitStatus == state.ERROR {
		return errors.New("mock state error")
	}
	if s.ExitStatus == state.PANIC {
		return errors.New("mock state panic")
	}
	return nil
}

func (s *MockState) Run(ctx context.Context) (context.Context, state.Status) {
	return ctx, s.ExitStatus
}

func (s *MockState) String() string {
	return "Mock"
}
