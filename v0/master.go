// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master.
package master

import (
	"github.com/munbot/master/v0/version"
)

// Version returns the running version information.
func Version() *version.Info {
	return new(version.Info)
}
