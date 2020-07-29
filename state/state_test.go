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
	count := int(lastStateID) - 1
	if len(stidMap) != count {
		t.Errorf("len stidMap(%d) != count(%d)", len(stidMap), count)
	}
}

func TestAllState(t *testing.T) {
	sm := NewMachine().(*SM)
	sm.stid = __zero
	for st := range stidMap {
		sm.newst = false
		if err := sm.SetState(st); err != nil {
			t.Fatalf("%v: %v", st, err)
		}
		if sm.State() != st {
			t.Errorf("sm.State() expect: %v - got: %v", st, sm.State())
		}
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
