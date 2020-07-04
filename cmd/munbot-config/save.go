// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/log"
)

func save(cfg *munbot.Config) {
	log.Debug("save...")
	if err := cfg.Save(); err != nil {
		log.Fatal("config save failed!")
	}
}
