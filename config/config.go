// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

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
	return nil
}
