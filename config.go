// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var fileOpen func(string) (*os.File, error) = os.Open

type Config struct {
	db map[string]config.Value
	reg map[string]*string
	Name string `json:"name,omitempty"`
	Test config.Value `json:"test,omitempty"`
	TestInt *config.IntValue `json:"testint,omitempty"`
	TestBool *config.BoolValue `json:"testbool,omitempty"`
}

func newConfig() *Config {
	return &Config{
		db: make(map[string]config.Value),
		reg: make(map[string]*string),
		Name: "munbot",
		Test: config.NewString("test", ""),
		TestInt: config.NewInt("testint", 0),
		TestBool: config.NewBool("testbool", false),
	}
}

func (c *Config) String() string {
	return c.Name
}

func (c *Config) register() {
	c.reg["name"] = &c.Name
	c.db["test"] = c.Test
	c.db["testint"] = c.TestInt
	c.db["testbool"] = c.TestBool
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
	cfg.register()
	return cfg
}

func (c *Config) Read(fh io.ReadCloser) error {
	defer fh.Close()
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return log.Error(err)
	}
	if err := json.Unmarshal(blob, c); err != nil {
		return log.Error(err)
	}
	return nil
}

//~ func (c *Config) Write(fh io.WriteCloser) error {
func (c *Config) Write(fh io.Writer) error {
	//~ defer fh.Close()
	blob, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return log.Error(err)
	}
	fh.Write(blob)
	fh.Write([]byte("\n"))
	//~ if err := ioutil.WriteFile(filename, blob, 0644); err != nil {
		//~ return log.Error(err)
	//~ }
	return nil
}

func (c *Config) Dump() {
	log.Debug("dump")
	for k, v := range c.reg {
		fmt.Printf("%s=%s\n", k, *v)
	}
	for k, v := range c.db {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func (c *Config) Update(key, newval string) error {
	if _, ok := c.reg[key]; !ok {
		return errors.New(fmt.Sprintf("invalid config key: %s", key))
	}
	*c.reg[key] = newval
	return nil
}
