// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

func newUser(cfg *config.Munbot, name string) error {
	log.Debugf("new name %s", name)
	if _, err := config.NewUser(cfg, name); err != nil {
		return err
	}
	return save(cfg)
}
