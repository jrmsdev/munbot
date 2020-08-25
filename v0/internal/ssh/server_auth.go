// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/session"
	"github.com/munbot/master/v0/internal/user"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/vfs"
)

var _ AuthManager = &ServerAuth{}

// AuthManager defines the server auth manager interface.
type AuthManager interface {
	Configure(dir string) error
	ServerConfig() *ssh.ServerConfig
	Login(fp, uid, sid string) (session.Token, error)
	Logout(session.Token) error
}

// ServerAuth implemenst the ssh server auth manager.
type ServerAuth struct {
	enable   bool
	name     string
	dir      string // config dir file path
	priv     string // ssh priv key file path
	keys     string // authorized_keys file path
	id       ssh.Signer
	lastHash string
}

// NewServerAuth creates a new ServerAuth instance.
func NewServerAuth() *ServerAuth {
	return &ServerAuth{
		enable: env.GetBool("MBAUTH"),
		name:   "master",
	}
}

func (a *ServerAuth) keyfp(pk ssh.PublicKey) string {
	return ssh.FingerprintSHA256(pk)
}

func (a *ServerAuth) setup() error {
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
		log.Warnf("%s: file not found.", a.priv)
		a.id, err = a.sshNewKeys(a.priv)
	}
	if err != nil {
		return err
	}
	log.Printf("Auth %s %s", a.name, a.keyfp(a.id.PublicKey()))
	return nil
}

func (a *ServerAuth) parseAuthKeys(fp string) (user.ID, error) {
	log.Debug("parse authorized keys")
	blob, err := vfs.ReadFile(a.keys)
	if err != nil {
		return user.Nil, log.Error(err)
	}
	for len(blob) > 0 {
		key, info, _, rest, err := ssh.ParseAuthorizedKey(blob)
		if err != nil {
			return user.Nil, log.Error(err)
		}
		if fp == a.keyfp(key) {
			log.Debugf("valid key %s", fp)
			return user.Parse(info)
		}
		blob = rest
	}
	return user.Nil, log.Errorf("Auth key %s", fp)
}

func (a *ServerAuth) publicKeyCallback(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {
	fp := a.keyfp(k)
	uid, err := a.parseAuthKeys(fp)
	if err != nil {
		return nil, err
	}
	log.Debugf("allow uid: %s", uid)
	return &ssh.Permissions{Extensions: map[string]string{
		"pubkey-fp":     fp,
		"x-munbot-user": uid.String(),
	}}, nil
}

// Configure sets up the auth directory.
func (a *ServerAuth) Configure(dir string) error {
	var err error
	a.dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	log.Debugf("config dir: %s", a.dir)
	return a.setup()
}

// ServerConfig creates a new instance of ssh.ServerConfig based on our settings.
func (a *ServerAuth) ServerConfig() *ssh.ServerConfig {
	cfg := &ssh.ServerConfig{}
	cfg.ServerVersion = "SSH-2.0-Munbot"
	cfg.MaxAuthTries = 3
	if a.id == nil {
		return cfg
	}
	cfg.AddHostKey(a.id)
	if a.enable {
		cfg.PublicKeyCallback = a.publicKeyCallback
	} else {
		log.Warn("ssh authentication is disabled!")
		cfg.PublicKeyCallback = a.publicKeyDisabled
	}
	return cfg
}

func (a *ServerAuth) publicKeyDisabled(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {
	return nil, fmt.Errorf("ssh auth disabled")
}

func (a *ServerAuth) sshLoadKeys(fn string) (ssh.Signer, error) {
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

func (a *ServerAuth) sshKeygen(filename string) error {
	cmd := exec.Command("ssh-keygen", "-v", "-N", "", "-h", "-t", "ed25519",
		"-f", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug(cmd.String())
	return cmd.Run()
}

func (a *ServerAuth) sshNewKeys(fn string) (ssh.Signer, error) {
	log.Print("SSH server generate new keys.")
	log.Debugf("new keys: %s", fn)
	if err := a.sshKeygen(fn); err != nil {
		return nil, log.Error(err)
	}
	return a.sshLoadKeys(fn)
}

func (a *ServerAuth) Login(fp, uid, sid string) (session.Token, error) {
	log.Infof("Auth login %s %s", fp, sid)
	return session.FromString(sid)
}

func (a *ServerAuth) Logout(s session.Token) error {
	log.Infof("Auth logout %s", s.String())
	return nil
}
