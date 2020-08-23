// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh

import (
	"fmt"
	"io/ioutil"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/vfs"
)

var ErrConfigDir error = errors.New("auth: config dir not found")

var _ AuthManager = &ServerAuth{}

// AuthManager defines the server auth manager interface.
type AuthManager interface {
	Configure(dir string) error
	ServerConfig() *ssh.ServerConfig
	Login(fp, sid string) error
	Logout(sid string) error
	Session(sid string) (error, error)
}

// ServerAuth implemenst the ssh server auth manager.
type ServerAuth struct {
	enable   bool
	name     string
	dir      string
	priv     string
	keys     string
	id       ssh.Signer
	auth     map[string]bool
	rw       *sync.RWMutex
	lastHash string
}

// NewServerAuth creates a new ServerAuth instance.
func NewServerAuth() *ServerAuth {
	return &ServerAuth{
		enable: env.GetBool("MBAUTH"),
		name:   "master",
		auth:   map[string]bool{},
		rw:     new(sync.RWMutex),
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

func (a *ServerAuth) parseAuthKeys() error {
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
				log.Warnf("%s: file was removed", a.keys)
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

func (a *ServerAuth) publicKeyCallback(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {
	if err := a.parseAuthKeys(); err != nil {
		return nil, err
	}
	a.rw.RLock()
	defer a.rw.RUnlock()
	fp := a.keyfp(k)
	if a.auth[fp] {
		log.Debugf("valid key %q", fp)
		return &ssh.Permissions{Extensions: map[string]string{"pubkey-fp": fp}}, nil
	}
	return nil, log.Errorf("Auth key %s", fp)
}

// Configure sets up the auth directory.
func (a *ServerAuth) Configure(dir string) error {
	var err error
	a.dir, err = filepath.Abs(dir)
	if a.dir == "." {
		return ErrConfigDir
	}
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

func (a *ServerAuth) Login(fp, sid string) error {
	log.Infof("Auth login %s %s", fp, sid)
	return nil
}

func (a *ServerAuth) Logout(sid string) error {
	log.Infof("Auth logout %s", sid)
	return nil
}

func (a *ServerAuth) Session(sid string) (error, error) {
	log.Debugf("%s new session", sid)
	return nil, nil
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
