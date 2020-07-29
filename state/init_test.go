// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state_test

import (
	"context"
	"testing"

	"github.com/munbot/master/core"
	"github.com/munbot/master/state"

	"github.com/munbot/master/testing/mock/state/machine"
	"github.com/munbot/master/testing/suite"
)

type InitSuite struct {
	*suite.Suite
	rt     core.Runtime
	mocksm *machine.MockSM
	sm     state.Machine
	ctx    context.Context
}

func TestSuite(t *testing.T) {
	suite.Run(t, &InitSuite{Suite: suite.New()})
}

func (s *InitSuite) SetupTest() {
	s.mocksm = machine.NewMockSM()
	s.rt = s.mocksm.MockRuntime
	s.sm = s.mocksm
	s.ctx = context.Background()
}

func (s *InitSuite) TearDownTest() {
	s.rt = nil
	s.mocksm = nil
	s.sm = nil
	s.ctx = nil
}

func (s *InitSuite) TestNew() {
	st := state.NewInitState(s.sm)
	s.Equal(state.Init.String(), st.String())
	s.Nil(st.Error())
}

func (s *InitSuite) TestRun() {
	s.NoError(s.sm.Init(nil, nil))
	st := state.NewInitState(s.sm)
	_, rc := st.Run(s.ctx)
	s.NoError(st.Error())
	s.Equal(state.OK, rc)
}

func (s *InitSuite) TestRunPanic() {
	st := state.NewInitState(s.sm)
	_, rc := st.Run(s.ctx)
	s.Equal(state.PANIC, rc)
}
