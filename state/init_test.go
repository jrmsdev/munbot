// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state_test

import (
	"testing"

	"github.com/munbot/master/state"

	"github.com/munbot/master/testing/mock/state/machine"
	"github.com/munbot/master/testing/suite"
)

type InitSuite struct {
	*suite.Suite
	mocksm *machine.MockSM
	sm     state.Machine
}

func TestSuite(t *testing.T) {
	suite.Run(t, &InitSuite{Suite: suite.New()})
}

func (s *InitSuite) SetupTest() {
	s.mocksm = machine.NewMockSM()
	s.sm = s.mocksm
}

func (s *InitSuite) TestNew() {
	st := state.NewInitState(s.sm)
	s.Equal(state.Init.String(), st.String())
}
