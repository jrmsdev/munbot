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

func (s *MachineSuite) SetupTest() {
}

func (s *MachineSuite) TestNew() {
	sm := NewMachine().(*sm)
	s.IsType(&InitState{}, sm.init)
	s.IsType(&ConfigureState{}, sm.configure)
	s.IsType(&StartState{}, sm.start)
	s.IsType(sm.st, sm.init)
	s.Nil(sm.Config)
	s.Nil(sm.ConfigFlags)
	s.Nil(sm.CoreFlags)
	s.Nil(sm.Runtime)
}

func (s *MachineSuite) TestInit() {
	require := s.Require()
	sm := NewMachine().(*sm)
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := sm.Init(cfg, cfl)
	require.NoError(err)
	s.IsType(config.New(), sm.Config)
	s.Equal(cfg, sm.ConfigFlags)
	s.Equal(cfl, sm.CoreFlags)
}

func (s *MachineSuite) TestRunCtxDone() {
	require := s.Require()
	m := NewMachine()
	cfg := &config.Flags{}
	cfl := &core.Flags{}
	err := m.Init(cfg, cfl)
	require.NoError(err)
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	err = m.Run(ctx)
	s.EqualError(err, "context canceled")
}

func (s *MachineSuite) TestRunError() {
}
