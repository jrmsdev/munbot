// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

func save(cfg *config.Munbot) error {
	log.Debug("save...")
	return config.Save(cfg)
}
