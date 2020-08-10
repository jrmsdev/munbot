// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package env manages settings configurable from os.Environ and .env files.
//
// Gets config settings from os.Environ or from Defaults otherwise. It will also
// load .env files if present.
//
// Settings from .env file will be used to populate os.Environ at init() time.
//
// If MBENV_CONFIG is set as "" (empty) no more files will be loaded, which is
// the default behavior.
//
// But if file ${MBENV_CONFIG}/${MBENV}.env exists it will be loaded too.
//
// Every loaded file overrides the settings from previous one.
package env

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gobuffalo/envy"

	"github.com/munbot/master/log"
)

// UNSET is the string returned for values not found in env nor in Defaults either.
const UNSET string = "__UNSET__"

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
	"MB_CONFIG":  filepath.FromSlash("/usr/local/etc/munbot"),

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

func init() {
	cfgdir := envy.Get("MBENV_CONFIG", MBENV_CONFIG)
	env := envy.Get("MBENV", MBENV)
	if cfgdir != "" {
		fn := filepath.Join(cfgdir, fmt.Sprintf("%s.env", env))
		envy.Load(fn)
	}
}

func defval(key string) string {
	v, ok := Defaults[key]
	if ok {
		return v
	}
	return UNSET
}

// Get key value, using Defaults for its default value. If not present, returns
// the "__UNSET__" string.
func Get(key string) string {
	return envy.Get(key, defval(key))
}

// GetBool returns the bool value for key.
// If there's a parsing error it will be logged and return default value false.
func GetBool(key string) bool {
	r, err := strconv.ParseBool(Get(key))
	if err != nil {
		log.Errorf("env parse bool %s: %s", key, err)
		return false
	}
	return r
}

// GetInt returns the int value for key.
// If there's a parsing error it will be logged and return default value 0.
func GetInt(key string) int {
	r, err := strconv.Atoi(Get(key))
	if err != nil {
		log.Errorf("env parse int %s: %s", key, err)
		return 0
	}
	return r
}

// GetUint returns the uint value for key.
// If there's a parsing error it will be logged and return default value 0.
func GetUint(key string) uint {
	r, err := strconv.ParseUint(Get(key), 10, 0)
	if err != nil {
		log.Errorf("env parse uint %s: %s", key, err)
		return 0
	}
	return uint(r)
}

// Set sets env key value. But it does not modify os.Environ.
func Set(key, val string) {
	envy.Set(key, val)
}

// SetInt sets an int value.
func SetInt(key string, val int) {
	envy.Set(key, strconv.FormatInt(int64(val), 10))
}

// SetUint sets an uint value.
func SetUint(key string, val uint) {
	envy.Set(key, strconv.FormatUint(uint64(val), 10))
}
