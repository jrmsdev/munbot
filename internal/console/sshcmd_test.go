// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/munbot/master/env"
	"github.com/munbot/master/internal/console"
	"github.com/munbot/master/log"
	"github.com/munbot/master/testing/suite"
)

func TestSSHCmdSuite(t *testing.T) {
	log.SetMode(env.Get("MB_LOG"))
	suite.Run(t, &sshCmdSuite{Suite: suite.New()})
}

type sshCmdSuite struct {
	*suite.Suite
	cmd    string
	skip   string
	cfg    *console.Config
	cons   *console.Console
	tmpdir string
	addr   string
	port   string
	ident  string
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
	if tmpdir, err := ioutil.TempDir("", "munbot_test_console_sshcmd_"); err != nil {
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
	s.port = s.cons.Addr().Port()
	s.T().Logf("setup server: %s", s.addr)
	s.ident = filepath.Join(s.tmpdir, "id_ed25519")
	cmd := exec.Command("ssh-keygen", "-q", "-f", s.ident,
		"-t", "ed25519", "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("setup ident %s", s.ident)
	authkeys := filepath.Join(s.tmpdir, "etc", "testing", "auth", "authorized_keys")
	if src, err := ioutil.ReadFile(s.ident + ".pub"); err != nil {
		s.T().Fatal(err)
	} else {
		if err := ioutil.WriteFile(authkeys, src, 0600); err != nil {
			s.T().Fatal(err)
		}
	}
	s.T().Logf("setup %s", authkeys)
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

type sshcmdTest struct {
	Args     []string
	ExitCode int
	Output   string
}

var allTests map[string]*sshcmdTest = map[string]*sshcmdTest{
	"Connect": {[]string{}, 255, "master> "},
}

func (s *sshCmdSuite) TestAll() {
	check := s.Require()
	check.NotNil(s.cons)
	for tname, tcmd := range allTests {
		buf := new(bytes.Buffer)
		cmd := exec.Command("ssh", "-p", s.port, "-i", s.ident, "-n", "-tt",
			"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
			"-F", filepath.FromSlash("./testdata/ssh_config"),
			"testing.munbot.local")
		cmd.Stdout = buf
		cmd.Stderr = buf
		cmd.Run()
		st := cmd.ProcessState
		check.True(st.Exited(), tname)
		check.Equal(tcmd.ExitCode, st.ExitCode(), tname)
		check.Equal(tcmd.Output, buf.String(), tname)
	}
}
