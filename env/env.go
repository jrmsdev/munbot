// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package env manages settings configurable from os.Environ and .env files.
package env

import (
	"fmt"
	"path/filepath"

	"github.com/gobuffalo/envy"
)

var Defaults map[string]string = map[string]string{
	"MUNBOT":         "munbot",
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

func Get(key string) string {
	return envy.Get(key, defval(key))
}
