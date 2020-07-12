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

type marshalFunc func(interface{}) ([]byte, error)

type Config struct {
	Munbot *Munbot `json:"munbot,omitempty"`
	marshal marshalFunc
}

func New() *Config {
	return &Config{Munbot: &Munbot{}, marshal: json.Marshal}
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
	return c.Read(fh)
}

func (c *Config) Read(r io.Reader) error {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(blob, c)
}

func (c *Config) Save(p *profile.Profile) error {
	fn := p.GetConfigFile()
	fh, err := vfs.Create(fn)
	if err != nil {
		return err
	}
	defer fh.Close()
	return c.Write(fh)
}

func (c *Config) Write(w io.Writer) error {
	blob, err := c.marshal(c)
	if err != nil {
		return err
	}
	if _, err := w.Write(blob); err != nil {
		return err
	}
	return nil
}
