// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package auth

import (
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/log"
)

func (a *Auth) sshLoadKeys() (ssh.Signer, error) {
	log.Debug("load keys")
	return nil, nil
}

func (a *Auth) sshKeygen(filename string) error {
	cmd := exec.Command("ssh-keygen", "-v", "-N", "", "-h", "-f", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug(cmd.String())
	return cmd.Run()
}

func (a *Auth) sshNewKeys() (ssh.Signer, error) {
	log.Debug("new keys")
	if err := a.sshKeygen(a.priv); err != nil {
		if err == exec.ErrNotFound {
			log.Warn(err)
			return nil, nil
		} else {
			return nil, log.Error(err)
		}
	}
	return a.sshLoadKeys()
}
