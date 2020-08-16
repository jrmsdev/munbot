// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package profile handles profiled settings.
package profile

import (
	"path/filepath"

	"github.com/munbot/master/env"
)

// Profile holds the profile settings.
type Profile struct {
	Name       string
	Home       string
	Config     string
	ConfigFile string
}

// New creates a new object with defaults values set.
func New() *Profile {
	return &Profile{
		Name:       env.Get("MB_PROFILE"),
		Home:       env.Get("MB_HOME"),
		Config:     env.Get("MB_CONFIG"),
		ConfigFile: "config.json",
	}
}

// String returns profile base path. But just for info/debug purposes. It will
// always be slash (/) separated.
func (p *Profile) String() string {
	d := filepath.Clean(p.Config)
	return filepath.ToSlash(filepath.Join(d, p.Name))
}

// GetPath returns the absolute profile named filepath.
func (p *Profile) GetPath(name string) string {
	d := filepath.Clean(p.Config)
	return filepath.Join(d, p.Name, name)
}

// GetConfigFile returns the absolute filename of the configuration file for the
// current os user. This is the filename used to save configuration updates.
func (p *Profile) GetConfigFile() string {
	return p.GetPath(p.ConfigFile)
}

// ListConfigFiles returns a list of all the profiled filenames to read the
// configuration from.
func (p *Profile) ListConfigFiles() []string {
	d := filepath.Clean(p.Config)
	return []string{
		filepath.Join(d, p.ConfigFile),
		p.GetConfigFile(),
	}
}
