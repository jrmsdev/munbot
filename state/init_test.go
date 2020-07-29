// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state_test

import (
	"context"
	"testing"

	"github.com/munbot/master/core"
	"github.com/munbot/master/state"

	"github.com/munbot/master/testing/mock/core/runtime"
	"github.com/munbot/master/testing/mock/state/machine"
	"github.com/munbot/master/testing/suite"
)

type InitSuite struct {
	*suite.Suite
	mockrt *runtime.MockRuntime
	rt     core.Runtime
	mocksm *machine.MockSM
	sm     state.Machine
	ctx    context.Context
}

func TestInitSuite(t *testing.T) {
	suite.Run(t, &InitSuite{Suite: suite.New()})
}

func (s *InitSuite) SetupTest() {
	s.mocksm = machine.NewMockSM()
	s.sm = s.mocksm
	s.mockrt = s.mocksm.MockRuntime
	s.rt = s.mockrt
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
	f := func() { st.Run(s.ctx) }
	s.PanicsWithError("init: run Master.Init first", f)
}

func (s *InitSuite) TestRunCtxDone() {
	st := state.NewInitState(s.sm)
	ctx2, cancel := context.WithCancel(s.ctx)
	cancel()
	_, rc := st.Run(ctx2)
	s.NoError(st.Error())
	s.Equal(state.DONE, rc)
}

func (s *InitSuite) TestRunSetStateError() {
	require := s.Require()
	require.NoError(s.mocksm.Init(nil, nil))
	s.mocksm.WithSetStateError = true
	st := state.NewInitState(s.mocksm)
	_, rc := st.Run(s.ctx)
	s.EqualError(st.Error(), "mock set state error: Configure")
	s.Equal(state.ERROR, rc)
}

func (s *InitSuite) TestRunLockError() {
	require := s.Require()
	require.NoError(s.mocksm.Init(nil, nil))
	s.mockrt.WithLockError = true
	st := state.NewInitState(s.mocksm)
	_, rc := st.Run(s.ctx)
	s.EqualError(st.Error(), "Init: mock lock error")
	s.Equal(state.ERROR, rc)
}
