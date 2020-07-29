// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"
	"testing"
)

func TestStatusMap(t *testing.T) {
	if len(stMap) != int(lastStatus) {
		t.Errorf("len stMap(%d) != lastStatus(%d)", len(stMap), lastStatus)
	}
}

func TestStateIDMap(t *testing.T) {
	if len(stidMap) != int(lastStateID) {
		t.Errorf("len stidMap(%d) != lastStateID(%d)", len(stidMap), lastStateID)
	}
}

var _ State = &MockState{}

type MockState struct {
	ExitStatus Status
}

func newMockState() *MockState {
	return &MockState{ExitStatus: OK}
}

func (s *MockState) Error() error {
	if s.ExitStatus == ERROR {
		return errors.New("mock state error")
	}
	if s.ExitStatus == PANIC {
		return errors.New("mock state panic")
	}
	return nil
}

func (s *MockState) Run(ctx context.Context) (context.Context, Status) {
	return ctx, s.ExitStatus
}

func (s *MockState) String() string {
	return "Mock"
}
