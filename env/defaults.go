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
// If it's set as "" (empty string) no extra files will be loaded (besides .env).
// By default it's set at init() time to ${HOME}/env if we can get user's home dir.
var MBENV_CONFIG string = ""

// Init contains the initial settings. They are copied at init(), so direct
// modifications to this map doesn't have any effect.
var Init map[string]string = map[string]string{
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
var defvals map[string]string

func init() {
	defvals = make(map[string]string)
	for k, v := range Init {
		defvals[k] = v
	}
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
	homeDir = filepath.Clean(envy.Get("MB_HOME", homeDir))
	defvals["MB_HOME"] = homeDir

	if configDir == "" || configDirErr != nil {
		configDir = filepath.Join(homeDir, "config")
	} else {
		configDir = filepath.Join(configDir, "munbot")
	}
	configDir = filepath.Clean(envy.Get("MB_CONFIG", configDir))
	defvals["MB_CONFIG"] = configDir
	MBENV_CONFIG = filepath.Join(homeDir, "env")
}

func loadEnv() {
	cfgdir := envy.Get("MBENV_CONFIG", MBENV_CONFIG)
	if cfgdir != "" {
		cfgdir = filepath.Clean(cfgdir)
		env := envy.Get("MBENV", MBENV)
		fn := filepath.Join(cfgdir, fmt.Sprintf("%s.env", env))
		envy.Load(fn)
	}
}

func init() {
	homeDir, homeDirErr = os.UserHomeDir()
	configDir, configDirErr = os.UserConfigDir()
	userDefaults()
	loadEnv()
}
