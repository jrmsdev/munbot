// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	gol "log"

	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

func TestDefaults(t *testing.T) {
	require := require.New(t)
	require.Equal(1, cdepth, "call depth")
	require.Equal(false, debug, "debug")
	require.Equal(gol.Llongfile, debugFlags, "debug flags")
	require.Equal(true, verbose, "verbose")
	require.Equal(gol.Ldate|gol.Ltime|gol.Lmicroseconds, stdFlags, "default flags")
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{Suite: suite.New(), buf: new(bytes.Buffer)})
}

type Suite struct {
	*suite.Suite
	buf *bytes.Buffer
}

func (s *Suite) SetupTest() {
	s.buf.Reset()
	debug = false
	info = true
	verbose = true
	l.SetPrefix("")
	l.SetColors("off")
	l.SetOutput(s.buf)
	l.SetFlags(stdFlags)
	l.SetDebug(false)
}

func (s *Suite) TestSetQuiet() {
	SetQuiet()
	s.False(debug, "debug enabled")
	s.False(info, "info enabled")
	s.False(verbose, "verbose enabled")
}

func (s *Suite) TestSetDebug() {
	SetQuiet()
	s.False(debug, "debug enabled")
	SetDebug()
	s.True(debug, "debug disabled")
	s.True(info, "info disabled")
	s.True(verbose, "verbose disabled")
}

func (s *Suite) TestSetInfo() {
	SetQuiet()
	s.False(info, "info enabled")
	SetInfo()
	s.False(debug, "debug enabled")
	s.True(info, "info disabled")
	s.False(verbose, "verbose enabled")
}

func (s *Suite) TestSetVerbose() {
	SetQuiet()
	s.False(verbose, "verbose enabled")
	SetVerbose()
	s.False(debug, "debug enabled")
	s.True(info, "info disabled")
	s.True(verbose, "verbose disabled")
}

func (s *Suite) TestSetPrefix() {
	require := s.Require()
	require.Equal("", l.Prefix(), "logger init prefix")
	SetPrefix("testing")
	prefix := fmt.Sprintf("[testing:%d] ", os.Getpid())
	s.Equal(prefix, l.Prefix(), "logger set prefix")
}

func (s *Suite) TestPrint() {
	Print("test")
	s.Regexp("^\\d\\d\\d\\d/\\d\\d/\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d\\d\\d\\d test\n$", s.buf.String(), "print msg")

	s.buf.Reset()
	Printf("te%s", "st")
	s.Regexp("\\d\\d test\n$", s.buf.String(), "printf msg")

	s.buf.Reset()
	SetQuiet()
	Print("test")
	s.Equal("", s.buf.String(), "print worked in quiet mode")
}

func (s *Suite) TestDebug() {
	require := s.Require()
	require.Equal(false, debug, "debug disabled")

	Debug("test")
	s.Equal("", s.buf.String(), "debug worked even if disabled")

	s.buf.Reset()
	SetDebug()
	Debug("test")
	//~ s.Regexp("^\\d\\d\\d\\d/\\d\\d/\\d\\d \\d\\d:\\d\\d:\\d\\d\\.\\d\\d\\d\\d\\d\\d .*log_test\\.go.* test\n$", s.buf.String(), "debug msg")
	s.Regexp("^.*log_test\\.go.* test\n$", s.buf.String(), "debug msg")

	s.buf.Reset()
	Debugf("te%s", "st")
	//~ s.Regexp("\\d .*log_test\\.go.* test\n$", s.buf.String(), "debugf msg")
	s.Regexp(".*log_test\\.go.* test\n$", s.buf.String(), "debugf msg")
}

func (s *Suite) TestError() {
	Error("test")
	s.Regexp("\\d\\d \\[ERROR\\] test\n$", s.buf.String(), "error msg")

	s.buf.Reset()
	Errorf("te%s", "st")
	s.Regexp("\\d\\d \\[ERROR\\] test\n$", s.buf.String(), "errorf msg")

	s.buf.Reset()
	SetQuiet()
	Error("test")
	s.Regexp("\\d\\d \\[ERROR\\] test\n$", s.buf.String(),
		"error should work even if in quiet mode")
}

func (s *Suite) TestWarn() {
	Warn("test")
	s.Regexp("\\d\\d \\[WARNING\\] test\n$", s.buf.String(), "warn msg")

	s.buf.Reset()
	Warnf("te%s", "st")
	s.Regexp("\\d\\d \\[WARNING\\] test\n$", s.buf.String(), "warnf msg")

	s.buf.Reset()
	SetQuiet()
	Warn("test")
	s.Equal("", s.buf.String(), "warn worked in quiet mode")
}

func (s *Suite) TestInfo() {
	SetInfo()
	Info("test")
	s.Regexp("\\d\\d test\n$", s.buf.String(), "info msg")

	s.buf.Reset()
	Infof("te%s", "st")
	s.Regexp("\\d\\d test\n$", s.buf.String(), "infof msg")

	s.buf.Reset()
	SetQuiet()
	Info("test")
	s.Equal("", s.buf.String(), "info msg in quiet mode")
}

func (s *Suite) TestPanic() {
	p := func() {
		Panic("test")
	}
	s.PanicsWithError("test", p, "log panic")
	s.Regexp("\\d\\d \\[PANIC\\] test\n$", s.buf.String(), "panic msg")

	s.buf.Reset()
	pf := func() {
		Panicf("te%s", "st")
	}
	s.PanicsWithError("test", pf, "log panicf")
	s.Regexp("\\d\\d \\[PANIC\\] test\n$", s.buf.String(), "panicf msg")
}

func (s *Suite) TestFatal() {
	var ret int
	exit := func(st int) {
		ret = st
	}
	osExit = exit
	defer func() {
		osExit = os.Exit
	}()
	Fatal("test")
	s.Regexp("\\d\\d \\[FATAL\\] test\n$", s.buf.String(), "fatal msg")
	s.Equal(2, ret, "exit code status")

	ret = 0
	s.buf.Reset()
	Fatalf("te%s", "st")
	s.Regexp("\\d\\d \\[FATAL\\] test\n$", s.buf.String(), "fatalf msg")
	s.Equal(2, ret, "exit code status")
}
