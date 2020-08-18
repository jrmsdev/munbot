// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package auth implements master's authentication.
package auth

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/env"
	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

var ErrCADir error = errors.New("auth: invalid CA dir")

// Auth implemenst the console auth manager.
type Auth struct {
	name     string
	dir      string
	priv     string
	keys     string
	id       ssh.Signer
	auth     map[string]bool
	rw       *sync.RWMutex
	lastHash string
}

// New creates a new Auth instance.
func New() *Auth {
	return &Auth{
		name: "master",
		auth: map[string]bool{},
		rw:   new(sync.RWMutex),
	}
}

func (a *Auth) keyfp(pk ssh.PublicKey) string {
	return ssh.FingerprintSHA256(pk)
}

func (a *Auth) setup() error {
	log.Debug("setup")
	a.name = env.Get("MUNBOT")
	var err error
	if err = vfs.MkdirAll(a.dir); err != nil {
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
		log.Printf("Auth %s %s", a.name, a.keyfp(a.id.PublicKey()))
		if err := a.parseAuthKeys(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Auth) parseAuthKeys() error {
	log.Debug("parse authorized keys")
	hash, herr := vfs.StatHash(a.keys)
	if herr != nil {
		if os.IsNotExist(herr) {
			if a.lastHash == "" || a.lastHash == "__notfound__" {
				if a.lastHash == "" {
					// warn only once...
					log.Warn(herr)
					a.lastHash = "__notfound__"
				}
				return nil
			} else {
				log.Warnf("%s: was file removed", a.keys)
				log.Debug("delete loaded keys")
				a.rw.Lock()
				for fp := range a.auth {
					delete(a.auth, fp)
				}
				a.lastHash = ""
				a.rw.Unlock()
				return nil
			}
		}
		return log.Error(herr)
	}
	if hash == a.lastHash {
		return nil
	}
	log.Print("Auth load keys...")
	blob, err := vfs.ReadFile(a.keys)
	if err != nil {
		return log.Error(err)
	}
	a.rw.Lock()
	defer a.rw.Unlock()
	a.lastHash = hash
	for fp := range a.auth {
		delete(a.auth, fp)
	}
	for len(blob) > 0 {
		key, _, _, rest, err := ssh.ParseAuthorizedKey(blob)
		if err != nil {
			return log.Error(err)
		}
		blob = rest
		fp := a.keyfp(key)
		a.auth[fp] = true
		log.Printf("Auth key %s", fp)
	}
	return nil
}

func (a *Auth) publicKeyCallback(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {
	if err := a.parseAuthKeys(); err != nil {
		return nil, err
	}
	a.rw.RLock()
	defer a.rw.RUnlock()
	fp := a.keyfp(k)
	if a.auth[fp] {
		return &ssh.Permissions{Extensions: map[string]string{"pubkey-fp": fp}}, nil
	}
	return nil, log.Errorf("Auth key %s", fp)
}
