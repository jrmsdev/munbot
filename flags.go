// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"flag"
)

// Flags holds main flags settings.
type Flags struct {
	Debug bool
	Profile string
}

func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.Debug, "debug", false, "enable debug settings")
	fs.StringVar(&f.Profile, "profile", "default", "profile `name`")
}
