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

func (s *MockState) Run(context.Context) Status {
	return s.ExitStatus
}

func (s *MockState) String() string {
	return "Mock"
}
