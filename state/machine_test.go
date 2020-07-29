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
	s.IsType(sm.st, sm.init)
	s.IsType(&core.Core{}, sm.rt)
	s.Nil(sm.cfg)
	s.Nil(sm.cfgFlags)
	s.Nil(sm.coreFlags)
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
	f := func() {
		sm.Run(context.TODO())
	}
	s.PanicsWithError("mock state panic", f)
}
