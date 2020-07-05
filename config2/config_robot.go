// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config2

type Robot struct {
	Enable  bool   `json:"enable,omitempty"`
	Name    string `json:"name,omitempty"`
	AutoRun bool   `json:"autorun,omitempty"`
}
