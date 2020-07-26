// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

func (s *Suite) TestMap() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")

	p := NewParser(c)
	m := p.Map("test")
	s.Equal(map[string]string{"test.opt": "testing"}, m, "parse map")
}

func (s *Suite) TestUpdate() {
	fh := s.fs.Add("test/config.json")
	fh.WriteString(`{"test":{"opt":"testing"}}`)
	c := New()
	err := c.Load(s.profile)
	s.require.NoError(err, "load error")

	p := NewParser(c)
	err = p.Update("noopt", "val")
	s.Error(err, "update error")

	x := c.Section("test")
	s.Equal("testing", x.Get("opt"), "testing opt")

	err = p.Update("test.opt", "newval")
	s.require.NoError(err, "update error")
	s.Equal("newval", x.Get("opt"), "testing opt new val")
}

func (s *Suite) TestSet() {
	c := New()
	p := NewParser(c)

	err := p.Set("test.opt", "testing")
	s.require.NoError(err, "set error")

	err = p.Set("test.opt", "error")
	s.require.Error(err, "set exist error")
}

func (s *Suite) TestSetOrUpdate() {
	c := New()
	p := NewParser(c)
	p.SetOrUpdate("test.opt", "testing")

	x := c.Section("test")
	s.Equal("testing", x.Get("opt"), "testing opt")

	p.SetOrUpdate("test.opt", "newval")
	s.Equal("newval", x.Get("opt"), "testing opt new val")

	p.SetOrUpdate("test.newopt", "testing")
	s.Equal("testing", x.Get("newopt"), "testing new opt val")
}
