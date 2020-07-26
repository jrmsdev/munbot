// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"

	"github.com/munbot/master/log"
)

// Flags holds some configuration settings that can be overriden from cmd args.
type Flags struct {
	Debug     bool
	Quiet     bool
	Verbose   bool
	Profile   string
	Config    string
	ConfigDir string
}

// NewFlags creates a new Flags object and sets the flags to the provided handler.
func NewFlags(fs *flag.FlagSet) *Flags {
	f := &Flags{}
	fs.BoolVar(&f.Debug, "debug", false, "enable debug settings")
	fs.BoolVar(&f.Quiet, "q", false, "set quiet mode")
	fs.BoolVar(&f.Verbose, "v", false, "set verbose mode")
	fs.StringVar(&f.Profile, "profile", "default", "profile `name`")
	fs.StringVar(&f.Config, "config", "config.json", "config file `name`")
	fs.StringVar(&f.ConfigDir, "cfg.dir", "", "config dir `path`")
	return f
}

// Parse parses the flags.
func (f *Flags) Parse() error {
	if f.Debug {
		log.DebugEnable()
	}
	if f.Quiet {
		log.SetQuiet()
	}
	if f.Verbose {
		log.SetVerbose()
	}
	return nil
}
