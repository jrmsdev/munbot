// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"context"
	"errors"
	"testing"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/state"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/suite"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	v := Version()
	assert.Regexp(`^\d+\.\d+\.\d+$`, v.String())
}

var _ state.Machine = &MockStateMachine{}

type MockStateMachine struct {
	WithInitError bool
	WithRunError  bool
}

func (m *MockStateMachine) Init(cf *config.Flags, fl *core.Flags) error {
	if m.WithInitError {
		return errors.New("mock init error")
	}
	return nil
}

func (m *MockStateMachine) Run(ctx context.Context) error {
	if m.WithRunError {
		return errors.New("mock run error")
	}
	return nil
}

type MasterSuite struct {
	*suite.Suite
	m *Master
	sm *MockStateMachine
}

func TestSuite(t *testing.T) {
	suite.Run(t, &MasterSuite{Suite: suite.New()})
}

func (s *MasterSuite) SetupTest() {
	s.sm = nil
	s.sm = new(MockStateMachine)
	s.m = nil
	s.m = New()
	s.m.sm = s.sm
}

func (s *MasterSuite) TestInit() {
	s.NoError(s.m.Init(nil, nil))
}

func (s *MasterSuite) TestInitError() {
	s.sm.WithInitError = true
	s.Error(s.m.Init(nil, nil))
}

func (s *MasterSuite) TestRun() {
	s.NoError(s.m.Run())
}

func (s *MasterSuite) TestRunError() {
	s.sm.WithRunError = true
	s.Error(s.m.Run())
}
