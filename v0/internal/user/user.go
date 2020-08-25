// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package user implemets internal user management.
package user

import (
	"net/mail"
	"strings"

	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/hash"
)

type User struct {
	e  *mail.Address
	id string
}

func Parse(emailAddr string) (*User, error) {
	emailAddr = strings.TrimSpace(emailAddr)
	if emailAddr == "" {
		return nil, log.Error("no user info")
	}
	e, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return nil, log.Error(err)
	}
	e.Name = strings.TrimSpace(e.Name)
	e.Address = strings.TrimSpace(e.Address)
	if e.Address == "" {
		return nil, log.Error("invalid credentials")
	}
	if e.Name == "" {
		e.Name = strings.Split(e.Address, "@")[0]
	}
	return &User{e: e, id: hash.Sum(e.Address)}, nil
}

func (u *User) Name() string {
	return u.e.Name
}

func (u *User) ID() string {
	return u.id
}

func (u *User) String() string {
	return u.id
}
