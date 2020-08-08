// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

import (
	"github.com/munbot/master/auth"
)

var _ AuthManager = &auth.Auth{}

// AuthManager defines the interface for server's auth manager.
type AuthManager interface {
}

// Config is the server config.
type Config struct {
	Enable bool
	Addr   string
	Port   uint
	Auth   AuthManager
}

// Server implements the ssh console server.
type Server struct {
	enable bool
	auth   AuthManager
}

func New() *Server {
	return &Server{}
}

func (s *Server) Configure(cfg *Config) error {
	if cfg.Enable {
		s.enable = true
		// TODO: check cfg.Auth is not nil or get a default one
		s.auth = cfg.Auth
	}
	return nil
}
