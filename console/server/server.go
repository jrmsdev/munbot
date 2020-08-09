// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

import (
	"fmt"
	"net"
	"sync"
	"time"

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
	wg     *sync.WaitGroup
	wait   time.Duration
}

func New() *Server {
	return &Server{
		done: make(chan bool, 1),
		wg:   &sync.WaitGroup{},
		wait: 300 * time.Millisecond,
	}
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
	// listen
	s.ln, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Debugf("listen error: %v", err)
		return err
	}
	defer s.ln.Close()
	log.Printf("Console server ssh://%s", s.addr)
	// accept connections
	fail := make(chan error, 1)
	s.wg.Add(1)
	go s.accept(fail)
	// monitor...
	defer close(fail)
LOOP:
	for {
		select {
		case err = <-fail:
			break LOOP
		case <-s.done:
			break LOOP
		default:
			time.Sleep(s.wait)
		}
	}
	s.wg.Wait()
	return err
}

func (s *Server) accept(fail chan<- error) {
	var err error
	defer s.wg.Done()
	// accept
	var nConn net.Conn
	nConn, err = s.ln.Accept()
	if err != nil {
		log.Debugf("accept error: %v", err)
		fail <- err
		return
	}
	defer nConn.Close()
	log.Printf("Console connected from %q", nConn.RemoteAddr())
	// ssh handshake
	conn, _, reqs, serr := ssh.NewServerConn(nConn, s.cfg)
	if serr != nil {
		log.Debugf("handshake error: %v", serr)
		fail <- serr
		return
	}
	defer conn.Close()
	log.Printf("Console handshake from %s@%s", conn.User(), conn.RemoteAddr())
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		ssh.DiscardRequests(reqs)
	}()
	log.Printf("Console login %s", conn.Permissions.Extensions["pubkey-fp"])
}

func (s *Server) Stop() error {
	defer close(s.done)
	err := s.ln.Close()
	s.wg.Wait()
	s.done <- true
	return err
}
