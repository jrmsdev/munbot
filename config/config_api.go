// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/internal/config"
)

type Api struct {
	Enable bool                `json:"enable,omitempty"`
	Addr   string              `json:"addr,omitempty"`
	Port   int                 `json:"port,omitempty"`
	Cert   *config.RelFilepath `json:"cert,omitempty"`
	Key    *config.RelFilepath `json:"key,omitempty"`
	Path   *config.AbsPath     `json:"path,omitempty"`
}
