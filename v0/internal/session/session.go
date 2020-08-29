// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package session implemets tools needed to manage user sessions.
package session

import (
	"sync"

	"github.com/munbot/master/v0/internal/user"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/uuid"
)

type Token uuid.UUID

var Nil Token = Token{}

func FromString(sid string) (Token, error) {
	var t Token
	if u, err := uuid.FromString(sid); err != nil {
		return Nil, log.Errorf("Session %s: %v", sid, err)
	} else {
		t = Token(u)
	}
	return t, nil
}

func (t Token) String() string {
	return uuid.ToString(uuid.UUID(t))
}

type Session struct {
	id  Token
	uid user.ID
	fp  string
}

type sm struct {
	*sync.Mutex
	sess map[Token]*Session
}

var sess *sm

func init() {
	sess = &sm{new(sync.Mutex), make(map[Token]*Session)}
}

func (m *sm) Add(sid Token, uid user.ID, fp string) (ok bool) {
	m.Lock()
	defer m.Unlock()
	if _, found := m.sess[sid]; found {
		return false
	}
	m.sess[sid] = &Session{sid, uid, fp}
	return true
}

func (m *sm) Remove(sid Token) (ok bool) {
	m.Lock()
	defer m.Unlock()
	if _, found := m.sess[sid]; found {
		return false
	}
	delete(m.sess, sid)
	return true
}

func Login(sid Token, uid user.ID, fp string) error {
	log.Debugf("%s login", sid)
	if !sess.Add(sid, uid, fp) {
		return log.Errorf("Session found %s.", sid)
	}
	log.Printf("User login %s %s.", uid, sid)
	return nil
}

func Logout(sid Token) error {
	log.Debugf("%s logout", sid)
	if !sess.Remove(sid) {
		return log.Errorf("Session not found %s.", sid)
	}
	log.Printf("User logout %s.", sid)
	return nil
}
