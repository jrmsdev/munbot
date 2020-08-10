// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobuffalo/envy"
)

// MBENV is the default env name.
var MBENV string = "life"

// MBENV_CONFIG if set, should be a dir path where to look for ${MBENV}.env file.
var MBENV_CONFIG string = ""

// Defaults contains the default settings.
var Defaults map[string]string = map[string]string{
	"MUNBOT": "master",

	"MB_LOG":     "verbose",
	"MB_DEBUG":   "false",
	"MB_PROFILE": "default",

	// these will be set at init() time based on os user env
	"MB_HOME":   "",
	"MB_CONFIG": "",

	"MBAPI":       "true",
	"MBAPI_DEBUG": "false",
	"MBAPI_ADDR":  "127.0.0.1",
	"MBAPI_PORT":  "6490",
	"MBAPI_PATH":  "/",

	"MBAUTH": "true",

	"MBCONSOLE":      "true",
	"MBCONSOLE_ADDR": "0.0.0.0",
	"MBCONSOLE_PORT": "6492",
}

var (
	homeDir      string
	homeDirErr   error
	configDir    string
	configDirErr error
)

func userDefaults() {
	if homeDir == "" || homeDirErr != nil {
		homeDir = filepath.FromSlash("./.munbot")
	} else {
		homeDir = filepath.Join(homeDir, ".munbot")
	}
	Defaults["MB_HOME"] = homeDir

	if configDir == "" || configDirErr != nil {
		configDir = filepath.Join(homeDir, "config")
	} else {
		configDir = filepath.Join(configDir, "munbot")
	}
	Defaults["MB_CONFIG"] = configDir
}

func init() {
	homeDir, homeDirErr = os.UserHomeDir()
	configDir, configDirErr = os.UserConfigDir()
	userDefaults()
	cfgdir := envy.Get("MBENV_CONFIG", MBENV_CONFIG)
	if cfgdir != "" {
		cfgdir = filepath.Clean(cfgdir)
		env := envy.Get("MBENV", MBENV)
		fn := filepath.Join(cfgdir, fmt.Sprintf("%s.env", env))
		envy.Load(fn)
	}
}
