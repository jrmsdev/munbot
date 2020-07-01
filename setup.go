// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"os"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func setup() {
	log.Debug("setup")
	if err := os.MkdirAll(flags.ConfigDir, 0770); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(flags.CacheDir, 0770); err != nil {
		log.Fatal(err)
	}
}
