// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's console server authentication.
package auth

import (
	"path/filepath"

	"github.com/munbot/master/log"
)

// Auth implemenst the console auth manager.
type Auth struct {
	dir string
}

// New creates a new Auth instance.
func New() *Auth {
	return &Auth{}
}

// Configure sets up the CA directory.
func (a *Auth) Configure(cadir string) error {
	a.dir = filepath.Clean(cadir)
	log.Debugf("CA dir: %s", a.dir)
	return nil
}
