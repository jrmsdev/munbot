// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/auth"
	"github.com/munbot/master/log"
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
	addr   string
	ln     net.Listener
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
		s.addr = fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
	}
	return nil
}

func (s *Server) Start() error {
	var err error
	s.ln, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Debugf("listen error: %v", err)
		return err
	}
	defer s.ln.Close()
	log.Printf("Console server ssh://%s", s.addr)
	var nConn net.Conn
	nConn, err = s.ln.Accept()
	if err != nil {
		log.Debugf("accept error: %v", err)
		return err
	}
	defer nConn.Close()
	log.Printf("Console connected from %q", nConn.RemoteAddr())
	<-s.done
	return nil
}

func (s *Server) Stop() error {
	s.done <- true
	return nil
}
