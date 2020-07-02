// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"os"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func SetupInfo() {
	log.Printf("Config dir %s", flags.ConfigDir)
	log.Printf("Cache dir %s", flags.CacheDir)
	log.Printf("Data dir %s", flags.DataDir)
}

func setup() {
	log.Debug("setup")
	if err := os.MkdirAll(flags.ConfigDir, 0770); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(flags.CacheDir, 0770); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(flags.DataDir, 0770); err != nil {
		log.Fatal(err)
	}
}
