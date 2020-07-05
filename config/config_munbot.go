// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Munbot struct {
	Master *Master          `json:"master,omitempty"`
	User   map[string]*User `json:"user,omitempty"`
}
