// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/log"
)

func newUser(cfg *munbot.Config, name string) {
	log.Debugf("new name %s", name)
	munbot.NewUser(cfg.NewUser(name))
	save(cfg)
}
