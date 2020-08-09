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
	env := Get("MBENV")
	if "test" == envy.Get("GO_ENV", env) {
		env = "test"
	}
	envy.Load(filepath.FromSlash(fmt.Sprintf("./env/%s.env", env)))
}

func defval(key string) string {
	return Defaults[key]
}

func Get(key string) string {
	return envy.Get(key, defval(key))
}
