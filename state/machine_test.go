// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"testing"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"

	"github.com/munbot/master/testing/suite"
)

type MachineSuite struct {
	*suite.Suite
	rt core.Runtime
	sm *SM
}

func TestMachineSuite(t *testing.T) {
	suite.Run(t, &MachineSuite{Suite: suite.New()})
}

func (s *MachineSuite) SetupTest() {
	s.rt = core.NewRuntime()
	s.sm = NewMachine(s.rt).(*SM)
}

func (s *MachineSuite) TearDownTest() {
	s.rt = nil
	s.sm = nil
}

func (s *MachineSuite) TestNew() {
	s.IsType(&InitState{}, s.sm.init)
	s.IsType(&ConfigureState{}, s.sm.configure)
	s.IsType(&StartState{}, s.sm.start)
	s.IsType(s.sm.init, s.sm.st)
	s.IsType(&core.Core{}, s.sm.rt)
	s.Equal(s.sm.rt, s.sm.Runtime())
	s.Nil(s.sm.cfg)
	s.Nil(s.sm.cfgFlags)
	s.Nil(s.sm.coreFlags)
	s.Equal(Init, s.sm.stid)
	s.Equal(Init, s.sm.State())
	s.True(s.sm.newst)
}

func (s *MachineSuite) TestSetState() {
	var err error
	s.sm.SetState(Init)
	s.Equal(Init, s.sm.State())
	err = s.sm.SetState(Init)
	s.Equal(ErrSetSameState, err)
	s.sm.SetState(Configure)
	s.Equal(Configure, s.sm.State())
	s.sm.SetState(Start)
	s.Equal(Start, s.sm.State())
	err = s.sm.SetState(-1)
	s.Equal(ErrSetInvalid, err)
}

func (s *MachineSuite) TestInit() {
	require := s.Require()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := s.sm.Init(cfg, cfl)
	require.NoError(err)
	s.IsType(config.New(), s.sm.Config())
	s.Equal(cfg, s.sm.ConfigFlags())
	s.Equal(cfl, s.sm.CoreFlags())
}

func (s *MachineSuite) TestRunCtxDone() {
	require := s.Require()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := s.sm.Init(cfg, cfl)
	require.NoError(err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = s.sm.Run(ctx)
	s.EqualError(err, "context canceled")
}

func (s *MachineSuite) TestRunExitNoNewState() {
	require := s.Require()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := s.sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = OK
	s.sm.st = st
	err = s.sm.Run(context.TODO())
	s.NoError(err)
}

func (s *MachineSuite) TestRunError() {
	require := s.Require()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := s.sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = ERROR
	s.sm.st = st
	s.sm.newst = true
	err = s.sm.Run(context.TODO())
	s.EqualError(err, "mock state error")
}

func (s *MachineSuite) TestRunPanic() {
	require := s.Require()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := s.sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = PANIC
	s.sm.st = st
	s.sm.newst = true
	f := func() {
		s.sm.Run(context.TODO())
	}
	s.PanicsWithError("mock state panic", f)
}
