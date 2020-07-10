// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"github.com/munbot/master/version"
)

func Version() *version.Info {
	return new(version.Info)
}
