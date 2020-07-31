// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/suite"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	v := Version()
	assert.Regexp(`^\d+\.\d+\.\d+$`, v.String())
}

type MasterSuite struct {
	*suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, &MasterSuite{Suite: suite.New()})
}

//~ func (s *MasterSuite) SetupTest() {
//~ }

//~ func (s *MasterSuite) TestInit() {
//~ }
