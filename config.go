// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

type Config struct {
	name string
}

func NewConfig(name string) *Config {
	return &Config{name}
}

func (c *Config) String() string {
	return c.name
}
