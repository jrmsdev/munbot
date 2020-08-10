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
	"MBENV": "life",
}

func init() {
	cfgdir := envy.Get("MBENV_CFGDIR", filepath.FromSlash("./env"))
	env := Get("MBENV")
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
