// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"context"
	"fmt"
	"io"
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
	log.Debug("stop")
	defer close(s.done)
	s.done <- true
	err := s.ln.Close()
	s.wg.Wait()
	return err
}

func (s *Console) Start() error {
	log.Debug("start")
	var err error
	// listen
	s.ln, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Debugf("listen error: %v", err)
		return err
	}
	log.Printf("Console server ssh://%s", s.addr)
	// accept connections
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.accept(ctx)
}

func (s *Console) accept(ctx context.Context) error {
	log.Debug("accept connections")
	var err error
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Debugf("context done, error: %v", err)
			return err
		case <-s.done:
			log.Debug("done!")
			return nil
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
			log.Debugf("%s serve context done: %v", sid, ctx.Err())
			return
		default:
		}
		t := nc.ChannelType()
		log.Debugf("%s serve channel type %s",sid, t)
		if t != "session" {
			nc.Reject(ssh.UnknownChannelType, "unknown channel type")
			log.Errorf("Console %s unknown channel type: %s", sid, t)
			continue
		}
		ch, req, err := nc.Accept()
		if err != nil {
			log.Errorf("Console %s could not accept channel: %v", sid, err)
			continue
		}
		s.wg.Add(1)
		go func(ctx context.Context, in <-chan *ssh.Request) {
			log.Debugf("%s serve request", sid)
			defer s.wg.Done()
			for req := range in {
				select {
				case <-ctx.Done():
					return
				default:
				}
				log.Debugf("%s serve request type %s", sid, req.Type)
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
			log.Debugf("%s serve shell", sid)
			defer s.wg.Done()
			defer ch.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				log.Debugf("%s shell read line...", sid)
				line, err := term.ReadLine()
				if err != nil {
					if err != io.EOF {
						log.Errorf("Console %s: %v", sid, err)
					}
					break
				}
				log.Printf("%s TERM: %s", sid, line)
			}
		}(ctx, ch)
	}
}
