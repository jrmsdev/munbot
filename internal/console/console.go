// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/internal/auth"
	"github.com/munbot/master/log"
)

// Config is the server config.
type Config struct {
	Enable bool
	Addr   string
	Port   uint
	Auth   auth.Manager
}

type Server interface {
	Configure(*Config) error
	Start() error
	Stop() error
}

var _ Server = &Console{}

// Console Consoleements the ssh console server.
type Console struct {
	enable bool
	auth   auth.Manager
	cfg    *ssh.ServerConfig
	done   chan bool
	addr   string
	ln     net.Listener
	wg     *sync.WaitGroup
	wait   time.Duration
}

func New() *Console {
	return &Console{
		done: make(chan bool, 1),
		wg:   &sync.WaitGroup{},
		wait: 300 * time.Millisecond,
	}
}

func (s *Console) Configure(cfg *Config) error {
	if cfg.Enable {
		s.enable = true
		// TODO: check cfg.Auth is not nil or get a default one
		s.auth = cfg.Auth
		s.cfg = s.auth.ServerConfig()
		s.addr = fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
	}
	return nil
}

func (s *Console) Stop() error {
	defer close(s.done)
	s.done <- true
	s.wg.Wait()
	return s.ln.Close()
}

func (s *Console) Start() error {
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
	ttl := time.Now().Add(time.Hour)
	ctx, cancel := context.WithDeadline(context.Background(), ttl)
	s.wg.Add(1)
	go s.accept(ctx)
	// monitor...
	defer cancel()
LOOP:
	for {
		select {
		case <-s.done:
			break LOOP
		default:
			time.Sleep(s.wait)
		}
	}
	return err
}

func (s *Console) accept(ctx context.Context) {
	var err error
	defer s.wg.Done()
	ttl, _ := ctx.Deadline()
LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Debugf("context done, error: %v", ctx.Err())
			break LOOP
		case <-s.done:
			break LOOP
		default:
		}
		var nc net.Conn
		nc, err = s.ln.Accept()
		if err != nil {
			log.Errorf("Console accept: %v", err)
			continue
		}
		if err := nc.SetDeadline(ttl); err != nil {
			log.Errorf("Console set connection deadline: %v", err)
			nc.Close()
			continue
		}
		s.wg.Add(1)
		go s.dispatch(ctx, nc)
	}
}

func (s *Console) dispatch(ctx context.Context, nc net.Conn) {
	defer s.wg.Done()
	defer nc.Close()
	select {
	case <-ctx.Done():
		return
	default:
	}
	// ssh handshake
	conn, _, reqs, err := ssh.NewServerConn(nc, s.cfg)
	if err != nil {
		log.Debugf("handshake error: %v", err)
		return
	}
	defer conn.Close()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		ssh.DiscardRequests(reqs)
	}()
	log.Printf("Auth login %s", conn.Permissions.Extensions["pubkey-fp"])
}
