// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"io"
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
		log.Debugf("read %s", fn)
	}
	return cfg
}
