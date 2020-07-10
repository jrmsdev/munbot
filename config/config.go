// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
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
	dirs []string
	Munbot *Munbot `json:"munbot,omitempty"`
}

func New(filename string, dirs ...string) *Config {
	return &Config{filename: filename, dirs: dirs}
}

func (c *Config) SetDefaults() {
	c.Munbot = &Munbot{
		Master: &Master{
			Enable: true,
			Name: "munbot",
		},
	}
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
		} else {
			return fmt.Errorf("%s: %s", name, err)
		}
	}
	defer fh.Close()
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}
	return json.Unmarshal(blob, c.Munbot)
}
