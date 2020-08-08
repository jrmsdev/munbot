// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's authentication.
package auth

import (
	"errors"
	"io/ioutil"
	"os"

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

func (a *Auth) setup() error {
	log.Debug("setup")
	var err error
	if err = os.MkdirAll(a.dir, 0700); err != nil {
		return err
	}
	if vfs.Exist(a.priv) {
		a.id, err = a.loadCA()
		log.Printf("Auth load CA %s", a.id.Recipient().String())
	} else {
		a.id, err = a.newCA()
		log.Printf("Auth new CA %s", a.id.Recipient().String())
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
	fh, err := vfs.Open(a.priv)
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	var blob []byte
	blob, err = ioutil.ReadAll(fh)
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	return age.ParseX25519Identity(string(blob))
}
