// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Config interface {
	Read() error
}

func New(filename string, dirs ...string) Config {
	return &Master{
		filename: filename,
		dirs: dirs,
	}
}

type Master struct {
	filename string
	dirs []string
}

func (c *Master) Read() error {
	return nil
}
