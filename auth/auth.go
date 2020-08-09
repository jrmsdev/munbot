// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's authentication.
package auth

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

var ErrCADir error = errors.New("auth: invalid CA dir")

// Auth implemenst the console auth manager.
type Auth struct {
	dir  string
	priv string
	pub  string
	id   ssh.Signer
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
	a.priv = filepath.Join(a.dir, "master_host")
	a.pub = filepath.Join(a.dir, "master_host.pub")
	if vfs.Exist(a.priv) {
		a.id, err = a.sshLoadKeys(a.priv)
		if err == nil && a.id != nil {
			log.Print("Auth loaded SSH keys")
		}
	} else {
		a.id, err = a.sshNewKeys(a.priv)
		if err == nil && a.id != nil {
			log.Print("Auth created SSH keys")
		}
	}
	if err != nil {
		return err
	}
	return nil
}
