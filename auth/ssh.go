// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package auth

import (
	"io/ioutil"
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

func (a *Auth) sshLoadKeys(fn string) (ssh.Signer, error) {
	log.Debugf("load keys: %s", fn)
	var pk ssh.Signer
	if fh, err := vfs.Open(fn); err != nil {
		return nil, log.Error(err)
	} else {
		defer fh.Close()
		if blob, err := ioutil.ReadAll(fh); err != nil {
			return nil, log.Error(err)
		} else {
			var err error
			if pk, err = ssh.ParsePrivateKey(blob); err != nil {
				return nil, log.Error(err)
			}
		}
	}
	return pk, nil
}

func (a *Auth) sshKeygen(filename string) error {
	cmd := exec.Command("ssh-keygen", "-v", "-N", "", "-h", "-t", "ed25519",
		"-f", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug(cmd.String())
	return cmd.Run()
}

func (a *Auth) sshNewKeys(fn string) (ssh.Signer, error) {
	log.Debugf("new keys: %s", fn)
	if err := a.sshKeygen(fn); err != nil {
		if err == exec.ErrNotFound {
			log.Warn(err)
			return nil, nil
		} else {
			return nil, log.Error(err)
		}
	}
	return a.sshLoadKeys(fn)
}
