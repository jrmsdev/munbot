// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/munbot/master/v0"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/testing/suite"
)

func TestSSHCmdSuite(t *testing.T) {
	log.SetMode(env.Get("MB_LOG"))
	suite.Run(t, &sshCmdSuite{Suite: suite.New(), m: master.New()})
}

type sshCmdSuite struct {
	*suite.Suite
	m      *master.Robot
	cmd    string
	skip   string
	tmpdir string
	addr   string
	port   string
	ident  string
}

func (s *sshCmdSuite) SetupSuite() {
	if s.cmd == "" {
		var err error
		s.cmd, err = exec.LookPath("ssh")
		if err != nil {
			s.skip = err.Error()
			s.T().Skip(s.skip)
		}
		if s.skip == "" {
			_, err = exec.LookPath("ssh-keygen")
			if err != nil {
				s.skip = err.Error()
				s.T().Skip(s.skip)
			}
		}
	} else if s.skip != "" {
		s.T().Skip(s.skip)
	}
	if tmpdir, err := ioutil.TempDir("", "munbot_test_sshcmd_"); err != nil {
		s.T().Fatal(err)
	} else {
		s.tmpdir = tmpdir
		env.Set("MB_CONFIG", filepath.Join(s.tmpdir, "etc"))
	}
	s.T().Log("run master robot")
	go func(t *testing.T, m *master.Robot) {
		if err := m.Run(); err != nil {
			t.Fatal(err)
		}
	}(s.T(), s.m)
	time.Sleep(time.Second)
	s.addr = s.m.Addr().String()
	s.T().Logf("sshd addr %q", s.addr)
	if s.addr == "ssh:" {
		s.T().Fatal("could not get sshd server address")
	}
	s.port = s.m.Addr().Port()
	s.T().Logf("sshd port: %q", s.port)
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

func (s *sshCmdSuite) TearDownSuite() {
	s.T().Log("stop master robot")
	if err := s.m.Stop(); err != nil {
		s.T().Log(err)
	}
	s.m = nil
	s.T().Logf("remove tmpdir %q", s.tmpdir)
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
	"Connect":   {[]string{}, 255, "logout\r\n"},
	"PtyReq":    {[]string{"-tt"}, 255, "logout\r\n"},
	"Tunnel":    {[]string{"-tt", "-L", "9999:localhost:6492"}, 255, "logout\r\n"},
	//~ "ExecError": {[]string{"testing"}, 255, ""},
}

func (s *sshCmdSuite) TestAll() {
	check := s.Require()
	check.NotNil(s.m)
	buf := new(bytes.Buffer)
	defer buf.Reset()
	for tname, tcmd := range allTests {
		buf.Reset()
		log.Printf("[TEST]: sshcmd/%s", tname)
		st := s.runCmd(buf, tcmd.Args)
		log.Print("[TEST]: *****")
		check.True(st.Exited(), tname)
		check.Equal(tcmd.ExitCode, st.ExitCode(), tname)
		if tcmd.Output == "" {
			check.Equal("", buf.String(), tname)
		} else {
			check.Contains(buf.String(), tcmd.Output, tname)
		}
	}
}

func (s *sshCmdSuite) runCmd(buf *bytes.Buffer, args []string) *os.ProcessState {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(cancel context.CancelFunc) {
		select {
		case <-time.After(2 * time.Second):
			cancel()
		}
	}(cancel)
	cmd := exec.CommandContext(ctx, "ssh", "-p", s.port, "-i", s.ident, "-n",
		"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
		"-F", filepath.FromSlash("./testdata/ssh_config"),
		"testing.munbot.local")
	cmd.Stdout = buf
	cmd.Stderr = buf
	for _, a := range args {
		cmd.Args = append(cmd.Args, a)
	}
	cmd.Run()
	return cmd.ProcessState
}

func (s *sshCmdSuite) TestSSHIdentError() {
	command, err := exec.LookPath("ssh")
	if err != nil {
		s.T().Skip(err)
	}
	buf := new(bytes.Buffer)
	defer buf.Reset()
	//~ cmd := exec.Command(command, "-p", s.port, "-n", "-i", s.ident,
	cmd := exec.Command(command, "-p", s.port, "-n",
		"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
		"-F", filepath.FromSlash("./testdata/ssh_config"),
		"testing.munbot.local")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err = cmd.Run()
	s.Error(err)
	st := cmd.ProcessState
	s.Equal(255, st.ExitCode())
	s.Equal("", buf.String())
}

func (s *sshCmdSuite) TestSSHCopyIDError() {
	command, err := exec.LookPath("ssh-copy-id")
	if err != nil {
		s.T().Skip(err)
	}
	buf := new(bytes.Buffer)
	defer buf.Reset()
	cmd := exec.Command(command, "-i", s.ident, "-p", s.port,
		"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
		"-o", "StrictHostKeyChecking=no",
		"127.0.0.1")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err = cmd.Run()
	s.Error(err)
	st := cmd.ProcessState
	s.Equal(1, st.ExitCode())
	s.Contains(buf.String(), "ERROR: exec request failed on channel 0")
}

func (s *sshCmdSuite) TestSCPError() {
	command, err := exec.LookPath("scp")
	if err != nil {
		s.T().Skip(err)
	}
	buf := new(bytes.Buffer)
	defer buf.Reset()
	cmd := exec.Command(command, "-i", s.ident, "-P", s.port,
		"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
		"-F", filepath.FromSlash("./testdata/ssh_config"),
		"/etc/passwd", "testing.munbot.local:/etc/passwd")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err = cmd.Run()
	s.Error(err)
	st := cmd.ProcessState
	s.Equal(1, st.ExitCode())
	s.Contains(buf.String(), "lost connection")
}

func (s *sshCmdSuite) TestSFTPError() {
	command, err := exec.LookPath("sftp")
	if err != nil {
		s.T().Skip(err)
	}
	buf := new(bytes.Buffer)
	defer buf.Reset()
	cmd := exec.Command(command, "-i", s.ident, "-P", s.port,
		"-o", fmt.Sprintf("UserKnownHostsFile=%s", os.DevNull),
		"-F", filepath.FromSlash("./testdata/ssh_config"),
		"testing.munbot.local")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err = cmd.Run()
	s.Error(err)
	st := cmd.ProcessState
	s.Equal(255, st.ExitCode())
	s.Contains(buf.String(), "Connection closed")
}

func (s *sshCmdSuite) TestSSHKeyScan() {
	command, err := exec.LookPath("ssh-keyscan")
	if err != nil {
		s.T().Skip(err)
	}
	buf := new(bytes.Buffer)
	defer buf.Reset()
	cmd := exec.Command(command, "-p", s.port, "127.0.0.1")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err = cmd.Run()
	s.NoError(err)
	st := cmd.ProcessState
	s.Equal(0, st.ExitCode())
	s.Contains(buf.String(), fmt.Sprintf("# 127.0.0.1:%s SSH-2.0-Munbot", s.port))
	fn := filepath.Join(s.tmpdir, "etc", "testing", "auth", "id_ed25519.pub")
	blob, ferr := ioutil.ReadFile(fn)
	if ferr != nil {
		s.T().Fatal(ferr)
	}
	fields := strings.Fields(string(blob))
	s.Len(fields, 3)
	s.Contains(buf.String(),
		fmt.Sprintf("[127.0.0.1]:%s %s %s", s.port, fields[0], fields[1]))
}
