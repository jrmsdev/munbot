// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	"github.com/munbot/master/profile"
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
	c := New()
	c.SetDefaults()
	require.Equal(defcfg, c.Munbot, "default config")
}

type Suite struct {
	*suite.Suite
	fs      *mock.Filesystem
	require *require.Assertions
	profile *profile.Profile
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{Suite: suite.New()})
}

func (s *Suite) SetupTest() {
	s.require = require.New(s.T())
	s.fs = mock.NewFilesystem("test/config.json")
	mock.SetFilesystem(s.fs)
	s.profile = profile.New("testing")
	s.profile.ConfigFilename = "config.json"
	s.profile.ConfigDir = "test"
	s.profile.ConfigSysDir = "sys"
	s.profile.ConfigDistDir = "dist"
}

func (s *Suite) TearDownTest() {
	s.require = nil
	s.fs = nil
	mock.SetDefaultFilesystem()
	s.profile = nil
}

func (s *Suite) TestLoad() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString("{}")
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "read error")
}

func (s *Suite) TestLoadError() {
	s.fs.WithReadError = true
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "mock read error", "read error")
}

func (s *Suite) TestJSONError() {
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "unexpected end of JSON input", "read error")
}

func (s *Suite) TestReadFileNotExist() {
	s.profile.ConfigFilename = "nofile.txt"
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "read file not exist")
}

func (s *Suite) TestOpenError() {
	s.fs.WithOpenError = true
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "mock open error", "open error")
}
