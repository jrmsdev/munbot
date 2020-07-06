// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"path/filepath"
)

var (
	Debug         bool   = false
	DebugApi      bool   = false
	Profile       string = "master"
	ConfigDir     string = filepath.FromSlash("./.munbot/config")
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir  string = filepath.FromSlash("/usr/local/etc/munbot")
	ConfigFile    string = "config.json"
	CacheDir      string = filepath.FromSlash("./.munbot/cache")
	DataDir       string = filepath.FromSlash("./.munbot")
	ApiAddr       string = "0.0.0.0"
	ApiPort       int    = 6492
	ApiCert       string = filepath.FromSlash("ssl/api/cert.pem")
	ApiKey        string = filepath.FromSlash("ssl/api/key.pem")
)
