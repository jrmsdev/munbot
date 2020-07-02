// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"github.com/jrmsdev/munbot/version"
)

func Version() *version.Info {
	return new(version.Info)
}
