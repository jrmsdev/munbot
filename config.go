// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

type MasterConfig struct {
	Name string `json:"name,omitempty"`
}

type Config struct {
	Master *MasterConfig `json:"master,omitempty"`
}

func newConfig() *Config {
	return &Config{
		&MasterConfig{Name: "munbot"},
	}
}

func (c *Config) String() string {
	return c.Master.Name
}

func (c *Config) Read(fh io.ReadCloser) error {
	defer fh.Close()
	return nil
}

var fileOpen = os.Open

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
				log.Debugf("read %s", err)
			} else {
				log.Panic(err)
			}
		}
		if err := cfg.Read(fh); err != nil {
			log.Debugf("read %s", fn)
			log.Panic(err)
		}
	}
	return cfg
}
