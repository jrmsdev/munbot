// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package profile handles profiled settings.
package profile

import (
	"path/filepath"
)

// Profile holds the profile settings.
type Profile struct {
	Name       string
	Home       string
	Config     string
	ConfigFile string
}

// New creates a new object with defaults values set.
func New(name string) *Profile {
	home := filepath.FromSlash("/var/local/munbot")
	return &Profile{
		Name:       name,
		Home:       home,
		Config:     filepath.Join(home, "config"),
		ConfigFile: "config.json",
	}
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
