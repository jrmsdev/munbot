// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"net/http"
)

type ServerConfig struct {
	Enable bool
	Addr   string
	Port   uint
}

type Server interface {
	Configure(*ServerConfig) error
	Start() error
	Stop() error
	Mount(path string, handler http.Handler)
}
