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

type ID string

var Nil ID = ID("")

func (i ID) String() string {
	return string(i)
}

func Parse(emailAddr string) (ID, error) {
	emailAddr = strings.TrimSpace(emailAddr)
	if emailAddr == "" {
		return Nil, log.Error("user: no info")
	}
	e, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return Nil, log.Error(err)
	}
	e.Name = strings.TrimSpace(e.Name)
	e.Address = strings.TrimSpace(e.Address)
	if e.Address == "" {
		return Nil, log.Error("user: invalid credentials")
	}
	if e.Name == "" {
		e.Name = strings.Split(e.Address, "@")[0]
	}
	return ID(hash.Sum(e.Address)), nil
}

type User struct {
	e  *mail.Address
	id ID
}

func (u *User) Name() string {
	return u.e.Name
}

func (u *User) ID() string {
	return u.id.String()
}

func (u *User) String() string {
	return u.id.String()
}
