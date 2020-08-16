// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package auth

import (
	"fmt"
	"path/filepath"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/log"
	"github.com/munbot/master/version"
)

var _ Manager = &Auth{}

type Manager interface {
	Configure(cadir string) error
	ServerConfig() *ssh.ServerConfig
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
	log.Debugf("CA dir: %s", a.dir)
	return a.setup()
}

// ServerConfig creates a new instance of ssh.ServerConfig based on our settings.
func (a *Auth) ServerConfig() *ssh.ServerConfig {
	cfg := &ssh.ServerConfig{}
	if a.id == nil {
		log.Warn("ssh authentication is disabled!")
		return cfg
	}
	cfg.AddHostKey(a.id)
	cfg.MaxAuthTries = 3
	cfg.PasswordCallback = a.passwordCallback
	cfg.KeyboardInteractiveCallback = a.keyboardCallback
	cfg.PublicKeyCallback = a.publicKeyCallback
	cfg.BannerCallback = a.bannerCallback
	return cfg
}

func (a *Auth) bannerCallback(conn ssh.ConnMetadata) string {
	return fmt.Sprintf("Munbot %s %s\n", a.name, version.String())
}

func (a *Auth) passwordCallback(c ssh.ConnMetadata, p []byte) (*ssh.Permissions, error) {
	return nil, log.Errorf("Password auth not supported")
}

func (a *Auth) keyboardCallback(c ssh.ConnMetadata, d ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
	return nil, log.Errorf("Password auth not supported")
}
