// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"path/filepath"
)

var (
	Debug         bool   = false
	DebugApi      bool   = false
	Name          string = "master"
	ConfigDir     string = filepath.FromSlash("~/.config/munbot")
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir  string = filepath.FromSlash("/usr/local/etc/munbot")
	ConfigFile    string = "config.json"
	CacheDir      string = filepath.FromSlash("~/.cache/munbot")
	DataDir       string = filepath.FromSlash("~/.munbot")
)
