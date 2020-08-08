// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"github.com/munbot/master/console/server"
)

var _ Server = &server.Server{}

// Server defines the ssh server interface.
type Server interface {
}

func NewServer() Server {
	return server.New()
}
