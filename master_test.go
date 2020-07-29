// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/mock/state/machine"
	"github.com/munbot/master/testing/suite"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	v := Version()
	assert.Regexp(`^\d+\.\d+\.\d+$`, v.String())
}

type MasterSuite struct {
	*suite.Suite
	m *Master
	sm *machine.MockSM
}

func TestSuite(t *testing.T) {
	suite.Run(t, &MasterSuite{Suite: suite.New()})
}

func (s *MasterSuite) SetupTest() {
	s.sm = nil
	s.sm = machine.NewMockSM()
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
	s.m.Init(nil, nil)
	s.NoError(s.m.Run())
}

func (s *MasterSuite) TestRunError() {
	s.sm.WithRunError = true
	s.Error(s.m.Run())
}
