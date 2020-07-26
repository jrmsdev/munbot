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
	f.Set(newTestFS())
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
}

func (s *Suite) TestFlagsParse() {
	f := new(Flags)
	f.Set(newTestFS())
	f.Parse()
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
}

func (s *Suite) TestFlagsDefaults() {
	f := new(Flags)
	f.Set(newTestFS())
	f.Parse()
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile, "default profile")
}
