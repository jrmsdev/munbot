// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's console server authentication.
package auth

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/munbot/master/log"
)

var ErrCADir error = errors.New("auth: invalid CA dir")

// Auth implemenst the console auth manager.
type Auth struct {
	dir  string
	priv string
	pub  string
}

// New creates a new Auth instance.
func New() *Auth {
	return &Auth{}
}

// Configure sets up the CA directory.
func (a *Auth) Configure(cadir string) error {
	var err error
	a.dir, err = filepath.Abs(cadir)
	if a.dir == "." {
		return ErrCADir
	}
	if err != nil {
		return err
	}
	if err = os.MkdirAll(a.dir, 0700); err != nil {
		return err
	}
	log.Debugf("CA dir: %s", a.dir)
	return a.setup()
}

func (a *Auth) setup() error {
	log.Debug("setup")
	return nil
}
