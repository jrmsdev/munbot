// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's console server authentication.
package auth

import (
	"errors"
	"os"
	"path/filepath"

	"filippo.io/age"

	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

var ErrCADir error = errors.New("auth: invalid CA dir")

// Auth implemenst the console auth manager.
type Auth struct {
	dir  string
	priv string
	pub  string
	id   *age.X25519Identity
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
	a.priv = filepath.Join(a.dir, "ca")
	a.pub = filepath.Join(a.dir, "ca.pub")
	return a.setup()
}

func (a *Auth) setup() error {
	log.Debug("setup")
	var err error
	if vfs.Exist(a.priv) {
		log.Print("Auth load CA")
		a.id, err = a.loadCA()
	} else {
		log.Print("Auth new CA")
		a.id, err = a.newCA()
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) newCA() (*age.X25519Identity, error) {
	id, err := age.GenerateX25519Identity()
	if err != nil {
		log.Debugf("age error: %v", err)
		return nil, err
	}
	if fh, err := vfs.Create(a.priv); err != nil {
		log.Debug(err)
		return nil, err
	} else {
		if _, err := fh.WriteString(id.String()); err != nil {
			log.Debug(err)
			return nil, err
		}
	}
	if fh, err := vfs.Create(a.pub); err != nil {
		log.Debug(err)
		return nil, err
	} else {
		if _, err := fh.WriteString(id.Recipient().String()); err != nil {
			log.Debug(err)
			return nil, err
		}
	}
	return id, nil
}

func (a *Auth) loadCA() (*age.X25519Identity, error) {
	return nil, nil
}
