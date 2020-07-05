// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"github.com/jrmsdev/munbot/config2"
	"github.com/jrmsdev/munbot/log"
)

type User struct {
	*config2.User
}

func NewUser(cfg *config2.User) *User {
	log.Debugf("new '%s'", cfg.Name)
	return &User{cfg}
}
