// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"errors"
	"testing"

	"github.com/munbot/master/config/internal/parser"
	"github.com/munbot/master/config/profile"
	"github.com/munbot/master/testing/mock"
	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

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
	__handler = nil
	__handler = parser.New()
}

func (s *Suite) TestLoad() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString("{}")
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "read error")
}

func (s *Suite) TestReadError() {
	s.fs.WithReadError = true
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "test/config.json: mock read error", "read error")
}

func (s *Suite) TestReadJSONError() {
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "test/config.json: unexpected end of JSON input", "read error")
}

func (s *Suite) TestLoadFileNotExist() {
	s.profile.ConfigFilename = "nofile.txt"
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "read file not exist")
}

func (s *Suite) TestOpenError() {
	s.fs.WithOpenError = true
	c := New()
	err := c.Load(s.profile)
	s.require.EqualError(err, "sys/config.json: mock open error", "open error")
}

func (s *Suite) TestLoadOverride() {
	c := New()
	c.SetDefaults(Defaults)
	//~ s.require.Equal("munbot", c.Munbot.Master.Name, "master name")

	// config file overrides system file
	sysfh := s.fs.Add("sys/config.json")
	sysfh.WriteString(`{"master":{"name":"sys"}}`)
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"master":{"name":"test"}}`)
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	//~ s.Equal("test", c.Munbot.Master.Name, "master name")

	// load system options if config file is empty (or not found)
	sysfh = s.fs.Add("sys/config.json")
	sysfh.WriteString(`{"master":{"name":"sys"}}`)
	fh = s.fs.Add("test/config.json")
	fh.WriteString(`{}`)
	err = c.Load(s.profile)
	s.require.NoError(err, "load error")
	//~ s.Equal("sys", c.Munbot.Master.Name, "master name")

	// dist config file overrides everything
	fh = s.fs.Add("test/config.json")
	fh.WriteString(`{"master":{"name":"test"}}`)
	sysfh = s.fs.Add("sys/config.json")
	sysfh.WriteString(`{"master":{"name":"sys"}}`)
	distfh := s.fs.Add("dist/config.json")
	distfh.WriteString(`{"master":{"name":"dist"}}`)
	err = c.Load(s.profile)
	s.require.NoError(err, "load error")
	//~ s.Equal("dist", c.Munbot.Master.Name, "master name")
}

func (s *Suite) TestSave() {
	s.fs.Add("test/testing/config.json")
	c := New()
	err := c.Save(s.profile)
	s.require.NoError(err, "save error")
}

func (s *Suite) TestSaveError() {
	c := New()
	err := c.Save(s.profile)
	s.require.EqualError(err, "stat test/testing/config.json.mock-notfound: no such file or directory", "save error")
}

func mockDump() ([]byte, error) {
	return nil, errors.New("mock dump error")
}

func (s *Suite) TestWriteDumpError() {
	s.fs.Add("test/testing/config.json")
	c := New()
	c.dump = mockDump
	err := c.Save(s.profile)
	s.require.EqualError(err, "mock dump error", "write dump error")
}

func (s *Suite) TestWriteError() {
	s.fs.WithWriteError = true
	s.fs.Add("test/testing/config.json")
	c := New()
	err := c.Save(s.profile)
	s.require.EqualError(err, "mock write error", "write error")
}
