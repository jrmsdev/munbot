// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
)

func newTestFS() *flag.FlagSet {
	return flag.NewFlagSet("testing", flag.PanicOnError)
}

func (s *Suite) TestFlagsSet() {
	f := new(Flags)
	f.set(newTestFS())
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
	s.Equal("", f.ApiAddr, "default api.addr")
}

func (s *Suite) TestFlagsParse() {
	f := new(Flags)
	f.set(newTestFS())
	c := New()
	f.parse(c)
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
	s.Equal("ECFGMISS:default.netaddr", f.ApiAddr, "default api.addr")
}

func (s *Suite) TestFlagsDefaults() {
	f := new(Flags)
	f.set(newTestFS())
	c := New()
	c.SetDefaults()
	f.parse(c)
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
	s.Equal("0.0.0.0", f.ApiAddr, "default api.addr")
	s.Equal(int(6492), f.ApiPort, "default api.port")
}
