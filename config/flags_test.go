// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
)

func newTestFS() *flag.FlagSet {
	return flag.NewFlagSet("testing", flag.PanicOnError)
}

func (s *Suite) TestFlagsDefaults() {
	f := NewFlags(newTestFS())
	s.False(f.Debug, "default debug")
	s.Equal("", f.Profile.Name, "default profile")
}

func (s *Suite) TestFlagsParse() {
	f := NewFlags(newTestFS())
	f.Parse()
	s.False(f.Debug, "default debug")
	s.Equal("default", f.Profile.Name, "default profile")
}
