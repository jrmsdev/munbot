// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/log"
)

func edit(cfg *munbot.Config, filter, args string) {
	log.Debug("edit...")
	if err := cfg.Update(filter, args); err != nil {
		log.Fatal(err)
	}
	save(cfg)
}
