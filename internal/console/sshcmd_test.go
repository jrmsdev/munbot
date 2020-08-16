// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/munbot/master/env"
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
	cons *console.Console
	tmpdir string
	addr string
}

func (s *sshCmdSuite) SetupTest() {
	if s.cmd == "" {
		var err error
		s.cmd, err = exec.LookPath("ssh")
		if err != nil {
			s.skip = "ssh command not found"
			s.T().Skip(s.skip)
		}
		if s.skip == "" {
			_, err = exec.LookPath("ssh-keygen")
			if err != nil {
				s.skip = "ssh-keygen command not found"
				s.T().Skip(s.skip)
			}
		}
	} else if s.skip != "" {
		s.T().Skip(s.skip)
	}
	if tmpdir, err := ioutil.TempDir("", "test_console_sshcmd_"); err != nil {
		s.T().Fatal(err)
	} else {
		s.tmpdir = tmpdir
		env.Set("MB_CONFIG", filepath.Join(s.tmpdir, "etc"))
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
	go func(t *testing.T, c console.Server) {
		if err := c.Start(); err != nil {
			t.Fatal(err)
		}
	}(s.T(), s.cons)
	time.Sleep(50 * time.Millisecond)
	s.addr = s.cons.Addr().String()
	if s.addr == "ssh:" {
		s.T().Fatal("could not get console server address")
	}
	s.T().Logf("setup done: %s", s.addr)
}

func (s *sshCmdSuite) TearDownTest() {
	if err := s.cons.Stop(); err != nil {
		s.T().Log(err)
	}
	s.cfg = nil
	s.cons = nil
	if err := os.RemoveAll(s.tmpdir); err != nil {
		s.T().Log(err)
	}
	s.tmpdir = ""
	env.Set("MB_CONFIG", "etc")
}

func (s *sshCmdSuite) TestStart() {
	check := s.Require()
	check.NotNil(s.cons)
}
