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
	done   chan bool
}

func New() *Server {
	return &Server{done: make(chan bool, 1)}
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

func (s *Server) Start() error {
	<-s.done
	return nil
}

func (s *Server) Stop() error {
	s.done <- true
	return nil
}
