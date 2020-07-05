// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Master struct {
	Name     string `json:"name,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Api      *Api   `json:"api,omitempty"`
	Robot    *Robot `json:"robot,omitempty"`
}
