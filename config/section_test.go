// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

func (s *Suite) TestSection() {
	c := New()
	s.require.False(c.HasSection("test"), "section test")
	s.require.False(c.HasOption("test", "opt"), "section test option")
	x := c.Section("test")
	s.require.Equal("default", x.Name(), "section default name")
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing"}}`)
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	s.True(c.HasSection("test"), "section test")
	s.True(c.HasOption("test", "opt"), "section test option")
	x = c.Section("test")
	s.require.Equal("test", x.Name(), "section test name")
}

func (s *Suite) TestSectionOption() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	x := c.Section("test")
	s.require.Equal("test", x.Name(), "section test name")
	s.require.True(x.HasOption("opt"), "test opt")
	s.Equal("testing", x.Get("opt"), "test opt value")
}

func (s *Suite) TestSectionGetBool() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing", "opt.bool":"true"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	x := c.Section("test")

	s.False(x.GetBool("opt"), "test opt bool error")
	s.True(x.GetBool("opt.bool"), "test opt bool")
}

func (s *Suite) TestSectionGetInt() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing", "opt.int":"128"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	x := c.Section("test")

	s.Equal(int(0), x.GetInt("opt"), "test opt int error")
	s.Equal(int(128), x.GetInt("opt.int"), "test opt int")
}

func (s *Suite) TestSectionGetUint() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing", "opt.uint":"128"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")
	x := c.Section("test")

	s.Equal(uint(0), x.GetUint("opt"), "test opt uint error")
	s.Equal(uint(128), x.GetUint("opt.uint"), "test opt uint")
}
