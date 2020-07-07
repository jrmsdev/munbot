// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"os"

	"github.com/munbot/master/config/flags"
	"github.com/munbot/master/log"
)

func setupInfo() {
	log.Printf("Config profile %s", flags.Profile)
	log.Printf("Config dir %s", flags.ConfigDir)
	log.Printf("Data dir %s", flags.DataDir)
	log.Printf("Cache dir %s", flags.CacheDir)
}

func setup() error {
	log.Debug("setup")
	var err error
	err = os.MkdirAll(flags.DataDir, 0770)
	if err != nil {
		return err
	}
	err = os.MkdirAll(flags.ConfigDir, 0770)
	if err != nil {
		return err
	}
	err = os.MkdirAll(flags.CacheDir, 0770)
	if err != nil {
		return err
	}
	return nil
}
