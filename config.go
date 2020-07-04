// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"io"
	"os"
	"path/filepath"

	cfg "github.com/jrmsdev/munbot/internal/config"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var fileOpen func(string) (*os.File, error) = os.Open

type Config struct {
	*cfg.Manager
	Master *config.Master `json:"master,omitempty"`
}

func newConfig() *Config {
	c := cfg.New()
	return &Config{c, config.NewMaster(c)}
}

func Configure() *Config {
	dirs := []string{
		flags.ConfigDistDir,
		flags.ConfigSysDir,
		flags.ConfigDir,
	}
	log.Debugf("configure %s %v", flags.ConfigFile, dirs)
	cfg := newConfig()
	for _, dn := range dirs {
		fn := filepath.Join(dn, flags.ConfigFile)
		fh, err := fileOpen(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				log.Panicf("%s: %s", fn, err)
			}
		} else {
			log.Debugf("read %s", fn)
			if err := cfg.Read(fh); err != nil {
				log.Panicf("%s: %s", fn, err)
			}
		}
	}
	setup()
	return cfg
}

func (c *Config) Read(fh io.ReadCloser) error {
	defer fh.Close()
	return c.Manager.Read(c, fh)
}

func (c *Config) Bytes() ([]byte, error) {
	return c.Manager.Bytes(c)
}

func (c *Config) Write(fh io.Writer) error {
	return c.Manager.Write(c, fh)
}

func (c *Config) NewUser(name string) *config.User {
	return config.NewUser(c.Master.Section, name)
}
