// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

func newUser(cfg *config.Munbot, name string) error {
	log.Debugf("new name %s", name)
	if cfg.User == nil {
		cfg.User = make(map[string]*config.User)
	}
	if _, err := config.NewUser(cfg, name); err != nil {
		return err
	}
	return save(cfg)
}
