// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	"github.com/munbot/master/testing/mock"
	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

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

type Suite struct {
	*suite.Suite
	fs *mock.Filesystem
	require *require.Assertions
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{suite.New(), nil, nil})
}

func (s *Suite) SetupTest() {
	s.require = require.New(s.T())
	s.fs = mock.NewFilesystem("config/testing.txt")
	mock.SetFilesystem(s.fs)
}

func (s *Suite) TearDownTest() {
	s.require = nil
	s.fs = nil
	mock.SetDefaultFilesystem()
}

func (s *Suite) TestRead() {
	fh := s.fs.Add("testing/config.json")
	fh.WriteString("{}")
	c := New("config.json", "testing")
	err := c.Read()
	s.require.NoError(err, "read error")
}

func (s *Suite) TestReadError() {
	s.fs.WithReadError = true
	c := New("testing.txt", "config")
	err := c.Read()
	s.require.EqualError(err, "mock read error", "read error")
}

func (s *Suite) TestJSONError() {
	fh := s.fs.Add("testing/config.json")
	fh.WriteString("{")
	c := New("config.json", "testing")
	err := c.Read()
	s.require.EqualError(err, "unexpected end of JSON input", "read error")
}

func (s *Suite) TestReadFileNotExist() {
	c := New("nofile.txt", "nodir")
	err := c.Read()
	s.require.NoError(err, "read file not exist")
}

func (s *Suite) TestOpenError() {
	s.fs.WithOpenError = true
	c := New("testing.txt", "config")
	err := c.Read()
	s.require.EqualError(err, "mock open error", "open error")
}
