// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"net/http"
)

type Server interface {
	Configure() error
	Start() error
	Stop() error
	Mount(path string, handler http.Handler)
}
