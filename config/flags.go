// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
	"path/filepath"

	"github.com/munbot/master/config/profile"
	"github.com/munbot/master/env"
	"github.com/munbot/master/log"
)

// Flags holds some configuration settings that can be overriden from cmd args.
type Flags struct {
	Quiet   bool
	Debug   bool
	Info    bool
	Verbose bool
	Name    string
	Profile *profile.Profile
}

// NewFlags creates a new Flags object and sets the flags to the provided handler.
func NewFlags(fs *flag.FlagSet) *Flags {
	f := &Flags{Profile: profile.New()}
	// log
	fs.BoolVar(&f.Quiet, "quiet",
		false, "quiet mode: errors only")
	fs.BoolVar(&f.Debug, "debug",
		false, "debug mode: all messages")
	fs.BoolVar(&f.Info, "info",
		false, "info mode: errors plus info")
	fs.BoolVar(&f.Verbose, "verbose",
		true, "verbose mode: all but debug messages")
	// profile
	fs.StringVar(&f.Name, "name", "", "master `robot` name")
	fs.StringVar(&f.Profile.Name, "profile", "", "config profile `name`")
	fs.StringVar(&f.Profile.Config, "config", "", "config `dir/path`")
	return f
}

// Parse parses the flags.
func (f *Flags) Parse() error {
	// log
	if f.Verbose {
		env.Set("MB_LOG", "verbose")
	}
	if f.Info {
		env.Set("MB_LOG", "info")
	}
	if f.Quiet {
		env.Set("MB_LOG", "quiet")
	}
	if f.Debug {
		env.Set("MB_LOG", "debug")
		env.Set("MB_DEBUG", "true")
	}
	log.SetMode(env.Get("MB_LOG"))
	// profile
	if f.Name == "" {
		f.Name = env.Get("MUNBOT")
	} else {
		env.Set("MUNBOT", f.Name)
	}
	if f.Profile.Name == "" {
		f.Profile.Name = env.Get("MB_PROFILE")
	} else {
		env.Set("MB_PROFILE", f.Profile.Name)
	}
	if f.Profile.Config == "" {
		f.Profile.Config = filepath.Clean(env.Get("MB_CONFIG"))
	} else {
		env.Set("MB_CONFIG", filepath.Clean(f.Profile.Config))
	}
	return nil
}
