// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package env manages settings configurable from os.Environ and .env files.
package env

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gobuffalo/envy"

	"github.com/munbot/master/log"
)

var Defaults map[string]string = map[string]string{
	"MUNBOT":         "master",
	"MB_LOG":         "quiet",
	"MB_DEBUG":       "false",
	"MB_CONFIG":      filepath.FromSlash("/usr/local/etc/munbot"),
	"MB_PROFILE":     "default",
	"MBAPI":          "true",
	"MBAPI_DEBUG":    "false",
	"MBAPI_ADDR":     "127.0.0.1",
	"MBAPI_PORT":     "6490",
	"MBCONSOLE":      "true",
	"MBCONSOLE_ADDR": "0.0.0.0",
	"MBCONSOLE_PORT": "6492",
}

func init() {
	cfgdir := envy.Get("MBENV_CONFIG", filepath.FromSlash("./env"))
	env := envy.Get("MBENV", "life")
	fn := filepath.Join(cfgdir, fmt.Sprintf("%s.env", env))
	envy.Load(fn)
}

func defval(key string) string {
	v, ok := Defaults[key]
	if ok {
		return v
	}
	return "__UNSET__"
}

// Get key value, using Defaults for its default value. If not present, returns
// the "__UNSET__" string.
func Get(key string) string {
	return envy.Get(key, defval(key))
}

// GetBool returns the bool value for key.
func GetBool(key string) bool {
	r, err := strconv.ParseBool(Get(key))
	if err != nil {
		log.Errorf("env parse bool %s: %s", key, err)
		return false
	}
	return r
}

// GetInt returns the int value for key.
func GetInt(key string) int {
	r, err := strconv.Atoi(Get(key))
	if err != nil {
		log.Errorf("env parse int %s: %s", key, err)
		return 0
	}
	return r
}

// GetUint returns the uint value for key.
func GetUint(key string) uint {
	r, err := strconv.ParseUint(Get(key), 10, 0)
	if err != nil {
		log.Errorf("env parse uint %s: %s", key, err)
		return 0
	}
	return uint(r)
}
