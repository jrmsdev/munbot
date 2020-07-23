// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
)

// Flags holds some configuration settings that can be overriden from cmd args.
type Flags struct {
	Debug bool
	Profile string
	ApiAddr string
}

func (f *Flags) set(fs *flag.FlagSet) {
	fs.BoolVar(&f.Debug, "debug", false, "enable debug settings")
	fs.StringVar(&f.Profile, "profile", "default", "profile `name`")
	fs.StringVar(&f.ApiAddr, "api.addr", "", "api network `address` to bind to")
}

func (f *Flags) parse(c *Config) {
	f.parseApi(c.Section("master.api"))
}

func (f *Flags) parseApi(s *Section) {
	if f.ApiAddr == "" {
		f.ApiAddr = s.Get("netaddr")
	}
}
