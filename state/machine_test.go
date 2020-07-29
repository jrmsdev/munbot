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
}

func TestSuite(t *testing.T) {
	suite.Run(t, &MachineSuite{Suite: suite.New()})
}

//~ func (s *MachineSuite) SetupTest() {
//~ }

func (s *MachineSuite) TestNew() {
	sm := NewMachine().(*SM)
	s.IsType(&InitState{}, sm.init)
	s.IsType(&ConfigureState{}, sm.configure)
	s.IsType(&StartState{}, sm.start)
	s.IsType(sm.init, sm.st)
	s.IsType(&core.Core{}, sm.rt)
	s.Equal(sm.rt, sm.Runtime())
	s.Nil(sm.cfg)
	s.Nil(sm.cfgFlags)
	s.Nil(sm.coreFlags)
	s.Equal(Init, sm.stid)
	s.Equal(Init, sm.State())
	s.True(sm.newst)
}

func (s *MachineSuite) TestSetState() {
	var err error
	sm := NewMachine().(*SM)
	sm.SetState(Init)
	s.Equal(Init, sm.State())
	err = sm.SetState(Init)
	s.Equal(ErrSetTwice, err)
	sm.newst = false
	err = sm.SetState(Init)
	s.Equal(ErrSetSameState, err)
	sm.newst = false
	sm.SetState(Configure)
	s.Equal(Configure, sm.State())
	sm.newst = false
	sm.SetState(Start)
	s.Equal(Start, sm.State())
	sm.newst = false
	err = sm.SetState(-1)
	s.Equal(ErrSetInvalid, err)
}

func (s *MachineSuite) TestInit() {
	require := s.Require()
	sm := NewMachine()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := sm.Init(cfg, cfl)
	require.NoError(err)
	s.IsType(config.New(), sm.Config())
	s.Equal(cfg, sm.ConfigFlags())
	s.Equal(cfl, sm.CoreFlags())
}

func (s *MachineSuite) TestRunCtxDone() {
	require := s.Require()
	m := NewMachine()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := m.Init(cfg, cfl)
	require.NoError(err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = m.Run(ctx)
	s.EqualError(err, "context canceled")
}

func (s *MachineSuite) TestRunExitNoNewState() {
	require := s.Require()
	sm := NewMachine().(*SM)
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = OK
	sm.st = st
	err = sm.Run(context.TODO())
	s.NoError(err)
}

func (s *MachineSuite) TestRunError() {
	require := s.Require()
	sm := NewMachine().(*SM)
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = ERROR
	sm.st = st
	sm.newst = true
	err = sm.Run(context.TODO())
	s.EqualError(err, "mock state error")
}

func (s *MachineSuite) TestRunPanic() {
	require := s.Require()
	sm := NewMachine().(*SM)
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := sm.Init(cfg, cfl)
	require.NoError(err)

	st := newMockState()
	st.ExitStatus = PANIC
	sm.st = st
	sm.newst = true
	f := func() {
		sm.Run(context.TODO())
	}
	s.PanicsWithError("mock state panic", f)
}
