// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package machine

import (
	"github.com/munbot/master/state"
)

var _ state.Machine = &MockSM{}

type MockSM struct {
	*state.SM
}

func NewMockSM() *MockSM {
	return &MockSM{SM: state.NewMachine().(*state.SM)}
}
