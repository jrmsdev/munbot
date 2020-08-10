// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"

	"github.com/munbot/master/config/profile"
	"github.com/munbot/master/env"
	"github.com/munbot/master/log"
)

// Flags holds some configuration settings that can be overriden from cmd args.
type Flags struct {
	Debug   bool
	Quiet   bool
	Verbose bool
	Profile *profile.Profile
}

// NewFlags creates a new Flags object and sets the flags to the provided handler.
func NewFlags(fs *flag.FlagSet) *Flags {
	f := &Flags{
		Profile: profile.New(env.Get("MB_PROFILE")),
	}
	fs.BoolVar(&f.Debug, "debug", false, "enable debug settings")
	fs.BoolVar(&f.Quiet, "q", false, "set quiet mode")
	fs.BoolVar(&f.Verbose, "v", false, "set verbose mode")
	fs.StringVar(&f.Profile.Name, "profile",
		f.Profile.Name, "profile `name`")
	fs.StringVar(&f.Profile.Config, "config",
		f.Profile.Config, "config file `name`")
	fs.StringVar(&f.Profile.ConfigDir, "cfg.dir",
		f.Profile.ConfigDir, "config dir `path`")
	fs.StringVar(&f.Profile.ConfigSysDir, "cfg.sysdir",
		f.Profile.ConfigSysDir, "system config dir `path`")
	fs.StringVar(&f.Profile.ConfigDistDir, "cfg.distdir",
		f.Profile.ConfigDistDir, "dist config dir `path`")
	return f
}

// Parse parses the flags.
func (f *Flags) Parse() error {
	loglvl := env.Get("MB_LOG")
	if f.Verbose || "verbose" == loglvl {
		log.SetVerbose()
	}
	if f.Quiet || "quiet" == loglvl {
		log.SetQuiet()
	}
	if f.Debug || env.GetBool("MB_DEBUG") || "debug" == loglvl {
		log.DebugEnable()
	}
	return nil
}
