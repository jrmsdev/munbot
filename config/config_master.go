// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Master struct {
	Name string `json:"name,omitempty"`
	Enable bool `json:"enable,omitempty"`
}
