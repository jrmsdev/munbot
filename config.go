// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

type MasterConfig struct {
	Name string `json:"name,omitempty"`
}

type Config struct {
	Master *MasterConfig `json:"master,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		&MasterConfig{Name: "munbot"},
	}
}

func (c *Config) String() string {
	return c.Master.Name
}
