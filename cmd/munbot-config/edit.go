// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func edit(cfg *munbot.Config, filter, args string) {
	log.Debug("edit...")
	if err := cfg.Update(filter, args); err != nil {
		log.Fatal(err)
	}
	fn := filepath.Join(flags.ConfigDir, flags.ConfigFile)
	blob, err := cfg.Bytes()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fn, blob, 0660); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s saved", fn)
}
