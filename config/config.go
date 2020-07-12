// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

type Munbot struct {
	Master *Master `json:"master,omitempty"`
}

type Config struct {
	filename string
	dirs     []string
	Munbot   *Munbot `json:"munbot,omitempty"`
}

func New(filename string, dirs ...string) *Config {
	return &Config{filename, dirs, &Munbot{}}
}

func (c *Config) SetDefaults() {
	c.Munbot.Master = &Master{Enable: true, Name: "munbot"}
}

func (c *Config) Read() error {
	for i := range c.dirs {
		dn := c.dirs[i]
		fn := filepath.Join(dn, c.filename)
		log.Debugf("config try %s", fn)
		if err := c.readFile(fn); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) readFile(name string) error {
	fh, err := vfs.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
			return nil
		} else {
			return err
		}
	}
	defer fh.Close()
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}
	return json.Unmarshal(blob, c.Munbot)
}
