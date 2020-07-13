// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"
	"strings"
)

type Section interface {
	Update(string, string) error
	List() []string
}

type Config struct {
	db map[string]Section
	idx map[string]string
}

func NewConfig() *Config {
	return &Config{db: make(map[string]Section), idx: make(map[string]string)}
}

func (c *Config) AddSection(name string, s Section) error {
	if _, ok := c.db[name]; ok {
		return fmt.Errorf("duplicate section: %s", name)
	}
	c.db[name] = s
	for _, n := range s.List() {
		opt := fmt.Sprintf("%s.%s", name, n)
		if found, ok := c.idx[opt]; ok {
			return fmt.Errorf("%s option already owned by %s section", opt, found)
		}
		c.idx[opt] = name
	}
	return nil
}

func (c *Config) Update(option, newval string) error {
	sn, ok := c.idx[option]
	if !ok {
		return fmt.Errorf("invalid option: %s", option)
	}
	s := c.db[sn]
	opt := strings.Replace(option, sn + ".", "", 1)
	return s.Update(opt, newval)
}
