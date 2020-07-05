// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config2

type Api struct {
	Enable bool   `json:"enable,omitempty"`
	Addr   string `json:"addr,omitempty"`
	Port   int    `json:"port,omitempty"`
	Cert   string `json:"cert,omitempty"`
	Key    string `json:"key,omitempty"`
	Path   string `json:"path,omitempty"`
}
