// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"os"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func setupInfo() {
	log.Printf("Config dir %s", flags.ConfigDir)
	log.Printf("Cache dir %s", flags.CacheDir)
	log.Printf("Data dir %s", flags.DataDir)
}

func setup() error {
	log.Debug("setup")
	err := os.MkdirAll(flags.ConfigDir, 0770)
	if err != nil {
		return err
	}
	err = os.MkdirAll(flags.CacheDir, 0770)
	if err != nil {
		return err
	}
	err = os.MkdirAll(flags.DataDir, 0770)
	if err != nil {
		return err
	}
	return nil
}
