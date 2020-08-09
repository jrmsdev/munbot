// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

import (
	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/auth"
)

// Config is the server config.
type Config struct {
	Enable bool
	Addr   string
	Port   uint
	Auth   auth.Manager
}

// Server implements the ssh console server.
type Server struct {
	enable bool
	auth   auth.Manager
	cfg    *ssh.ServerConfig
}

func New() *Server {
	return &Server{}
}

func (s *Server) Configure(cfg *Config) error {
	if cfg.Enable {
		s.enable = true
		// TODO: check cfg.Auth is not nil or get a default one
		s.auth = cfg.Auth
		s.cfg = s.auth.ServerConfig()
	}
	return nil
}
