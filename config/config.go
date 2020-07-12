// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/munbot/master/log"
	"github.com/munbot/master/profile"
	"github.com/munbot/master/vfs"
)

type Munbot struct {
	Master *Master `json:"master,omitempty"`
}

type Config struct {
	Munbot *Munbot `json:"munbot,omitempty"`
}

func New() *Config {
	return &Config{&Munbot{}}
}

func (c *Config) SetDefaults() {
	c.Munbot.Master = &Master{Enable: true, Name: "munbot"}
}

func (c *Config) Load(p *profile.Profile) error {
	for _, fn := range p.ListConfigFiles() {
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
	return json.Unmarshal(blob, c)
}

func (c *Config) Read(r io.Reader) error {
	return nil
}

func (c *Config) Save(p *profile.Profile) error {
	return nil
}

func (c *Config) Write(w io.Writer) error {
	//~ blob, err := json.Marshal(c.Munbot)
	return nil
}
