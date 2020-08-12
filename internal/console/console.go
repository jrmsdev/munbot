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
	"golang.org/x/crypto/ssh/terminal"

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
	err := s.ln.Close()
	s.wg.Wait()
	return err
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.wg.Add(1)
	go s.accept(ctx)
	// monitor...
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
	log.Debug("accept connections")
	var err error
	defer s.wg.Done()
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
		s.wg.Add(1)
		go s.dispatch(ctx, nc)
	}
}

func (s *Console) dispatch(ctx context.Context, nc net.Conn) {
	log.Debug("dispatch")
	defer s.wg.Done()
	defer nc.Close()
	select {
	case <-ctx.Done():
		return
	default:
	}
	// ssh handshake
	conn, chans, reqs, err := ssh.NewServerConn(nc, s.cfg)
	if err != nil {
		log.Debugf("handshake error: %v", err)
		return
	}
	defer conn.Close()
	s.wg.Add(1)
	go func(reqs <-chan *ssh.Request) {
		defer s.wg.Done()
		ssh.DiscardRequests(reqs)
	}(reqs)
	// ctx session
	ctx = s.ctxNewSession(ctx)
	sid := s.ctxSession(ctx)
	// serve
	fp := conn.Permissions.Extensions["pubkey-fp"]
	log.Printf("Auth login %s %s", fp, sid)
	s.serve(ctx, chans)
	log.Printf("Auth logout %s", sid)
}

func (s *Console) serve(ctx context.Context, chans <-chan ssh.NewChannel) {
	sid := s.ctxSession(ctx)
	log.Debugf("serve session %s", sid)
	for nc := range chans {
		select {
		case <-ctx.Done():
			log.Debugf("serve context done: %v", ctx.Err())
			return
		default:
		}
		t := nc.ChannelType()
		log.Debugf("serve channel type %s", t)
		if t != "session" {
			nc.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		ch, req, err := nc.Accept()
		if err != nil {
			log.Errorf("could not accept channel: %v", err)
			continue
		}
		s.wg.Add(1)
		go func(ctx context.Context, in <-chan *ssh.Request) {
			log.Debug("serve request")
			defer s.wg.Done()
			for req := range in {
				select {
				case <-ctx.Done():
					return
				default:
				}
				log.Debugf("serve request type %s", req.Type)
				serve := false
				switch req.Type {
				case "pty-req":
					serve = true
				case "shell":
					serve = true
				}
				req.Reply(serve, nil)
			}
		}(ctx, req)
		term := terminal.NewTerminal(ch, "munbot> ")
		s.wg.Add(1)
		go func(ctx context.Context, ch ssh.Channel) {
			log.Debug("serve shell")
			defer s.wg.Done()
			defer ch.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				log.Debug("shell read line...")
				line, err := term.ReadLine()
				if err != nil {
					log.Error(err)
					break
				}
				log.Printf("TERM: %s", line)
			}
		}(ctx, ch)
	}
}
