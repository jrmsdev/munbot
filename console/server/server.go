// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

import (
	"github.com/munbot/master/console/auth"
)

var _ AuthManager = &auth.Auth{}

// AuthManager defines the interface for server's auth manager.
type AuthManager interface {
	Configure(cadir string) error
}

// Config is the server config.
type Config struct {
	Enable bool
	Addr   string
	Port   uint
	CADir  string
}

// Server implements the ssh console server.
type Server struct {
	enable bool
	auth   AuthManager
}

func New() *Server {
	return &Server{auth: auth.New()}
}

func (s *Server) Configure(cfg *Config) error {
	if cfg.Enable {
		s.enable = true
		return s.auth.Configure(cfg.CADir)
	}
	return nil
}
