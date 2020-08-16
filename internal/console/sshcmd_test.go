// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console_test

import (
	"os/exec"
	"testing"

	"github.com/munbot/master/internal/console"
	"github.com/munbot/master/testing/suite"
)

func TestSSHCmdSuite(t *testing.T) {
	suite.Run(t, &sshCmdSuite{Suite: suite.New()})
}

type sshCmdSuite struct {
	*suite.Suite
	cmd  string
	skip string
	cfg  *console.Config
	cons console.Server
}

func (s *sshCmdSuite) SetupTest() {
	if s.cmd == "" {
		var err error
		s.cmd, err = exec.LookPath("ssh")
		if err != nil {
			s.skip = "ssh command not found"
			s.T().Skip(s.skip)
		}
	} else if s.skip != "" {
		s.T().Skip(s.skip)
	}
	s.cfg = &console.Config{
		Enable: true,
		Addr:   "127.0.0.1",
		Port:   0,
	}
	s.cons = console.New()
	if err := s.cons.Configure(s.cfg); err != nil {
		s.T().Fatal(err)
	}
	s.T().Log("setup done")
}

func (s *sshCmdSuite) TearDownTest() {
	s.cfg = nil
	s.cons = nil
}

func (s *sshCmdSuite) TestStart() {
	check := s.Require()
	check.NotNil(s.cons)
}
