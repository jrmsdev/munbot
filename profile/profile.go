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
	configDirErr   error
	Name           string
	ConfigFilename string
	ConfigDir      string
	ConfigSysDir   string
	ConfigDistDir  string
)

func init() {
	homeDir, homeDirErr = os.UserHomeDir()
	ConfigDir, configDirErr = os.UserConfigDir()
	setDefaults()
}

func setDefaults() {
	if homeDir == "" || homeDirErr != nil {
		homeDir = filepath.FromSlash("./.munbot")
	} else {
		homeDir = filepath.Join(homeDir, ".munbot")
	}
	if ConfigDir == "" || configDirErr != nil {
		ConfigDir = filepath.Join(homeDir, "config")
	} else {
		ConfigDir = filepath.Join(ConfigDir, "munbot")
	}
	Name = "munbot"
	ConfigFilename = "config.json"
	ConfigSysDir = filepath.FromSlash("/usr/local/etc/munbot")
	ConfigDistDir = filepath.FromSlash("/etc/munbot")
}
