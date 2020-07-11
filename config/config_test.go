// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"
	"testing"

	"github.com/munbot/master/testing/mock"
	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

var tdir string = filepath.FromSlash("./testdata")
var defcfg *Munbot = &Munbot{
	Master: &Master{
		Enable: true,
		Name:   "munbot",
	},
}

func TestDefaults(t *testing.T) {
	require := require.New(t)
	c := New("empty")
	c.SetDefaults()
	require.Equal(defcfg, c.Munbot, "default config")
}

func TestRead(t *testing.T) {
	c := New("config.json", tdir)
	c.Read()
}

type Suite struct {
	*suite.Suite
	fs *mock.Filesystem
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{suite.New(), nil})
}

func (s *Suite) SetupTest() {
	s.fs = mock.NewFilesystem("config/testing.txt")
	mock.SetFilesystem(s.fs)
}

func (s *Suite) TearDownTest() {
	s.fs = nil
	mock.SetDefaultFilesystem()
}

func (s *Suite) TestReadError() {
	s.fs.WithReadError = true
	require := require.New(s.T())
	c := New("testing.txt", "config")
	err := c.Read()
	require.EqualError(err, "mock read error", "read error")
}

func (s *Suite) TestJSONError() {
	require := require.New(s.T())
	fh := s.fs.Add("testing/config.json")
	fh.WriteString("{")
	c := New("config.json", "testing")
	err := c.Read()
	require.EqualError(err, "unexpected end of JSON input", "read error")
}
