// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/internal/config"
)

type Master struct {
	*config.Section
	Name *config.StringValue `json:"name,omitempty"`
}
