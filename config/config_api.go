// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"

	"github.com/jrmsdev/munbot/internal/config"
)

type Api struct {
	Enable *config.BoolValue     `json:"enable,omitempty"`
	Addr   *config.StringValue   `json:"addr,omitempty"`
	Port   *config.IntValue      `json:"port,omitempty"`
	Cert   *config.FilepathValue `json:"cert,omitempty"`
	Key    *config.FilepathValue `json:"key,omitempty"`
	Path   *config.PathValue     `json:"path,omitempty"`
}

func newApi(s *config.Section, enable bool) *Api {
	return &Api{
		Enable: s.NewBool("api.enable", enable),
		Addr:   s.NewString("api.addr", "0.0.0.0"),
		Port:   s.NewInt("api.port", 3000),
		Cert:   s.NewFilepath("api.cert", filepath.FromSlash("ssl/api/cert.pem")),
		Key:    s.NewFilepath("api.key", filepath.FromSlash("ssl/api/key.pem")),
		Path:   s.NewPath("api.path", "/api"),
	}
}
