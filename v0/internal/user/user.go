// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package user implemets internal user management.
package user

import (
	"encoding/json"
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

func parse(emailAddr string) (*mail.Address, error) {
	emailAddr = strings.TrimSpace(emailAddr)
	if emailAddr == "" {
		return nil, log.Error("user: no info")
	}
	e, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return nil, log.Error(err)
	}
	e.Name = strings.TrimSpace(e.Name)
	e.Address = strings.TrimSpace(e.Address)
	if e.Address == "" {
		return nil, log.Error("user: invalid credentials")
	}
	if e.Name == "" {
		e.Name = strings.Split(e.Address, "@")[0]
	}
	return e, nil
}

type jsonUser struct {
	EName string
	EAddr string
	FP    string
	ID    string
}

func (j *jsonUser) String() string {
	b, err := json.Marshal(j)
	if err != nil {
		log.Panic(err)
		return ""
	}
	return string(b)
}

func Marshal(emailAddr, fp string) (string, error) {
	e, err := parse(emailAddr)
	if err != nil {
		return "", err
	}
	u := New(e, fp)
	var b []byte
	b, err = json.Marshal(&jsonUser{u.Name(), u.e.Address, u.Fingerprint(), u.String()})
	if err != nil {
		return "", log.Error(err)
	}
	return string(b), nil
}

type User struct {
	e  *mail.Address
	fp string
	id ID
}

func New(e *mail.Address, fp string) *User {
	return &User{e: e, fp: fp, id: ID(hash.Sum(e.Address))}
}

func Unmarshal(s string) (*User, error) {
	var j jsonUser
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return nil, log.Error(err)
	}
	e := &mail.Address{
		Name:    j.EName,
		Address: j.EAddr,
	}
	u := New(e, j.FP)
	if u.String() != j.ID {
		return nil, log.Error("user: invalid credentials")
	}
	return u, nil
}

func (u *User) Name() string {
	return u.e.Name
}

func (u *User) ID() ID {
	return u.id
}

func (u *User) Fingerprint() string {
	return u.fp
}

func (u *User) String() string {
	return u.id.String()
}
