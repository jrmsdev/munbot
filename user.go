// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

type User struct {
	*config.User
}

func NewUser(cfg *config.User) *User {
	log.Debugf("new '%s'", cfg.Name)
	return &User{cfg}
}
