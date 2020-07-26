// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
)

// Flags holds some configuration settings that can be overriden from cmd args.
type Flags struct {
	Debug   bool
	Profile string
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.Debug, "debug", false, "enable debug settings")
	fs.StringVar(&f.Profile, "profile", "default", "profile `name`")
}

// Parse parses the flags.
func (f *Flags) Parse() error {
	return nil
}
