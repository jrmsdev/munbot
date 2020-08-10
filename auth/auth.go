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
	keys string
	id   ssh.Signer
	auth map[string]bool
}

// New creates a new Auth instance.
func New() *Auth {
	return &Auth{auth: map[string]bool{}}
}

func (a *Auth) keyfp(pk ssh.PublicKey) string {
	return ssh.FingerprintSHA256(pk)
}

func (a *Auth) setup() error {
	log.Debug("setup")
	var err error
	if err = os.MkdirAll(a.dir, 0750); err != nil {
		return log.Error(err)
	}
	a.priv = filepath.Join(a.dir, "id_ed25519")
	a.keys = filepath.Join(a.dir, "authorized_keys")
	if vfs.Exist(a.priv) {
		a.id, err = a.sshLoadKeys(a.priv)
	} else {
		a.id, err = a.sshNewKeys(a.priv)
	}
	if err != nil {
		return err
	}
	if a.id != nil {
		log.Printf("Auth master %s", a.keyfp(a.id.PublicKey()))
		if err := a.parseAuthKeys(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Auth) parseAuthKeys() error {
	log.Debug("parse authorized keys")
	if !vfs.Exist(a.keys) {
		log.Warnf("%s: file not found", a.keys)
		return nil
	}
	blob, err := vfs.ReadFile(a.keys)
	if err != nil {
		return err
	}
	for len(blob) > 0 {
		key, _, _, rest, err := ssh.ParseAuthorizedKey(blob)
		if err != nil {
			log.Debug(err)
			return err
		}
		blob = rest
		fp := a.keyfp(key)
		a.auth[fp] = true
		log.Printf("Auth added %s", fp)
	}
	return nil
}

func (a *Auth) publicKeyCallback(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {
	fp := a.keyfp(k)
	if a.auth[fp] {
		return &ssh.Permissions{Extensions: map[string]string{"pubkey-fp": fp}}, nil
	}
	return nil, log.Errorf("Auth key %s", fp)
}
