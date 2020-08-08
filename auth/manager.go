// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package auth

import (
	"path/filepath"

	"github.com/munbot/master/log"
)

type Manager interface {
	Configure(cadir string) error
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
	a.priv = filepath.Join(a.dir, "ca")
	a.pub = filepath.Join(a.dir, "ca.pub")
	return a.setup()
}
