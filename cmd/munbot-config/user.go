// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

func newUser(cfg *config.Munbot, name string) error {
	log.Debugf("new name %s", name)
	if _, err := config.NewUser(cfg, name); err != nil {
		return err
	}
	return save(cfg)
}
