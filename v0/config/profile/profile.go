// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package profile handles profiled settings.
package profile

import (
	"path/filepath"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/vfs"
)

// Profile holds the profile settings.
type Profile struct {
	Name       string
	Home       string
	Config     string
	ConfigFile string
	Run        string
}

// New creates a new object with defaults values set.
func New() *Profile {
	return &Profile{
		Name:       env.Get("MB_PROFILE"),
		Home:       filepath.Clean(env.Get("MB_HOME")),
		Config:     filepath.Clean(env.Get("MB_CONFIG")),
		ConfigFile: "config.json",
		Run:        filepath.Clean(env.Get("MB_RUN")),
	}
}

// String returns profile base path. But just for info/debug purposes. It will
// always be slash (/) separated.
func (p *Profile) String() string {
	return filepath.ToSlash(filepath.Join(p.Config, p.Name))
}

// GetPath returns the absolute profile named filepath.
func (p *Profile) GetPath(name string) string {
	return filepath.Join(p.Config, p.Name, name)
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

// GetRundir returns the profile named rundir path.
func (p *Profile) GetRundir() string {
	return filepath.Join(p.Run, p.Name)
}

// GetRundirPath returns an absolute rundir profile named filepath.
func (p *Profile) GetRundirPath(name string) string {
	return filepath.Join(p.GetRundir(), name)
}

// Setup checks runtime setup requirements.
func (p *Profile) Setup() error {
	log.Debug("setup...")
	mkdirs := map[string]string{
		"config": filepath.Join(p.Config, p.Name),
		"run":    filepath.Join(p.Run, p.Name),
	}
	for name, dir := range mkdirs {
		log.Debugf("setup %s dir %q", name, dir)
		if err := vfs.MkdirAll(dir); err != nil {
			return log.Error(err)
		}
	}
	return nil
}
