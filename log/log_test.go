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
	require.Equal(2, cdepth, "call depth")
	require.Equal(false, debug, "debug")
	require.Equal(gol.Ldate|gol.Ltime|gol.Lmicroseconds|gol.Llongfile,
		debugFlags, "debug flags")
	require.Equal(true, verbose, "verbose")
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{Suite: suite.New(), buf: new(bytes.Buffer)})
}

type Suite struct {
	*suite.Suite
	logger *gol.Logger
	buf    *bytes.Buffer
}

func (s *Suite) SetupTest() {
	s.buf.Reset()
	s.logger = nil
	s.logger = gol.New(s.buf, "", gol.LstdFlags)
	output = s.logger.Output
	setFlags = s.logger.SetFlags
	setPrefix = s.logger.SetPrefix
	debug = false
	verbose = true
}

func (s *Suite) TestDebugEnable() {
	require := s.Require()
	require.Equal(false, debug, "debug disabled")
	require.Equal(gol.LstdFlags, s.logger.Flags(), "logger init flags")
	DebugEnable()
	s.Equal(true, debug, "debug enabled")
	s.Equal(debugFlags, s.logger.Flags(), "logger debug flags")
	// check SetQuiet does nothing in debug mode
	SetQuiet()
	s.Equal(true, verbose, "verbose disabled by SetQuiet in debug mode")
}

func (s *Suite) TestSetQuiet() {
	require := s.Require()
	require.Equal(false, debug, "debug disabled")
	require.Equal(true, verbose, "verbose enabled")
	SetQuiet()
	s.Equal(false, verbose, "verbose enabled after SetQuiet")
}

func (s *Suite) TestSetVerbose() {
	require := s.Require()
	require.Equal(false, debug, "debug disabled")
	require.Equal(true, verbose, "verbose enabled")
	SetQuiet()
	s.Equal(false, verbose, "verbose enabled after SetQuiet")
	SetVerbose()
	s.Equal(true, verbose, "verbose disabled after SetVerbose")
}

func (s *Suite) TestSetPrefix() {
	require := s.Require()
	require.Equal("", s.logger.Prefix(), "logger init prefix")
	SetPrefix("testing")
	prefix := fmt.Sprintf("[testing:%d] ", os.Getpid())
	s.Equal(prefix, s.logger.Prefix(), "logger set prefix")
}

func (s *Suite) TestPrint() {
	Print("test")
	s.Regexp("^\\d\\d\\d\\d/\\d\\d/\\d\\d \\d\\d:\\d\\d:\\d\\d test\n$", s.buf.String(), "print msg")

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
	DebugEnable()
	Debug("test")
	s.Regexp("^\\d\\d\\d\\d/\\d\\d/\\d\\d \\d\\d:\\d\\d:\\d\\d\\.\\d\\d\\d\\d\\d\\d .*log_test\\.go.* test\n$", s.buf.String(), "debug msg")

	s.buf.Reset()
	Debugf("te%s", "st")
	s.Regexp("\\d .*log_test\\.go.* test\n$", s.buf.String(), "debugf msg")
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
