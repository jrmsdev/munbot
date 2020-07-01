// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Master struct {
	*Section
	Name *StringValue `json:"name,omitempty"`
}
