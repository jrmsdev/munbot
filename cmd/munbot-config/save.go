// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

func save(cfg *config.Munbot) error {
	log.Debug("save...")
	return config.Save(cfg)
}
