// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package profile

import (
	"os"
	"path/filepath"
)

var (
	homeDir        string
	homeDirErr     error
	configDir      string
	configDirErr   error
)

func init() {
	homeDir, homeDirErr = os.UserHomeDir()
	configDir, configDirErr = os.UserConfigDir()
}

type Profile struct {
	Name           string
	ConfigFilename string
	ConfigDir      string
	ConfigSysDir   string
	ConfigDistDir  string
}

func setDefaults(p *Profile) *Profile {
	if homeDir == "" || homeDirErr != nil {
		homeDir = filepath.FromSlash("./.munbot")
	} else {
		homeDir = filepath.Join(homeDir, ".munbot")
	}
	if configDir == "" || configDirErr != nil {
		configDir = filepath.Join(homeDir, "config")
	} else {
		configDir = filepath.Join(configDir, "munbot")
	}
	p.ConfigFilename = "config.json"
	p.ConfigDir = configDir
	p.ConfigSysDir = filepath.FromSlash("/usr/local/etc/munbot")
	p.ConfigDistDir = filepath.FromSlash("/etc/munbot")
	return p
}

func New(name string) *Profile {
	return setDefaults(&Profile{Name: name})
}

func (p *Profile) ListConfigFiles() []string {
	l := make([]string, 0)
	if p.ConfigSysDir != "" {
		p.ConfigSysDir = filepath.Clean(p.ConfigSysDir)
		l = append(l, filepath.Join(p.ConfigSysDir, p.ConfigFilename))
		l = append(l, filepath.Join(p.ConfigSysDir, p.Name, p.ConfigFilename))
	}
	p.ConfigDir = filepath.Clean(p.ConfigDir)
	l = append(l, filepath.Join(p.ConfigDir, p.ConfigFilename))
	l = append(l, filepath.Join(p.ConfigDir, p.Name, p.ConfigFilename))
	if p.ConfigDistDir != "" {
		p.ConfigDistDir = filepath.Clean(p.ConfigDistDir)
		l = append(l, filepath.Join(p.ConfigDistDir, p.ConfigFilename))
		l = append(l, filepath.Join(p.ConfigDistDir, p.Name, p.ConfigFilename))
	}
	return l
}
