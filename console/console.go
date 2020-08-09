// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"github.com/munbot/master/auth"
	"github.com/munbot/master/console/server"
	"github.com/munbot/master/core/flags"
)

var _ Server = &server.Server{}

// Server defines the ssh server interface.
type Server interface {
	Configure(*server.Config) error
	Start() error
	Stop() error
}

// NewServer creates a new Server instance.
func NewServer() Server {
	return server.New()
}

// Configure parses provided settings and calls server.Configure.
func Configure(s Server, kfl *flags.Flags, m auth.Manager) error {
	return s.Configure(&server.Config{
		Enable: kfl.ConsoleEnable,
		Addr:   kfl.ConsoleAddr,
		Port:   kfl.ConsolePort,
		Auth:   m,
	})
}
