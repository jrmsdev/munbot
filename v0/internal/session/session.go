// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package session implemets tools needed to manage user sessions.
package session

import (
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/uuid"
)

type Token uuid.UUID

var Nil Token = Token{}

func (t Token) String() string {
	return uuid.ToString(uuid.UUID(t))
}

func New(uid, sid string) (Token, error) {
	var t Token
	if u, err := uuid.FromString(sid); err != nil {
		return Nil, err
	} else {
		t = Token(u)
	}
	log.Printf("User %s %s", uid, t)
	return t, nil
}
