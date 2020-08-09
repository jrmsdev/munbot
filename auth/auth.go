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
	ssh  ssh.Signer
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
		a.ssh, err = a.sshLoadKeys()
		if err == nil && a.ssh != nil {
			log.Printf("Auth loaded SSH keys %s", a.priv)
		}
	} else {
		a.ssh, err = a.sshNewKeys()
		if err == nil && a.ssh != nil {
			log.Printf("Auth created SSH keys %s", a.priv)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
