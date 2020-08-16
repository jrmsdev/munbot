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

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/munbot/master/env"
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
	lock   *sync.Mutex
	q      map[string]net.Conn
	closed bool
}

func New() *Console {
	return &Console{
		done: make(chan bool, 1),
		wg:   &sync.WaitGroup{},
		lock: new(sync.Mutex),
		q:    make(map[string]net.Conn),
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
	if s.enable {
		defer close(s.done)
		s.done <- true
		s.closed = true
		var err error
		if s.ln != nil {
			err = s.ln.Close()
		}
		s.wg.Wait()
		return err
	}
	log.Debug("api server is disabled")
	return nil
}

func (s *Console) Start() error {
	log.Debug("start")
	if s.enable {
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
	log.Warn("Console server is disabled")
	return nil
}

func (s *Console) accept(ctx context.Context) error {
	log.Debug("accept connections")
	var err error
LOOP:
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Debugf("context done, error: %v", err)
			return err
		case <-s.done:
			log.Debug("done!")
			break LOOP
		default:
		}
		var nc net.Conn
		nc, err = s.ln.Accept()
		if err != nil {
			if !s.closed {
				log.Errorf("Console accept: %v", err)
			}
			continue
		}
		// ctx session
		var sid string
		ctx, sid = s.ctxNewSession(ctx)
		// dispatch
		s.wg.Add(1)
		go func(ctx context.Context, nc net.Conn, sid string) {
			defer s.wg.Done()
			defer nc.Close()
			s.lock.Lock()
			log.Debugf("%s queue", sid)
			s.q[sid] = nc
			s.lock.Unlock()
			s.dispatch(ctx, nc, sid)
			s.lock.Lock()
			log.Debugf("%s dequeue", sid)
			delete(s.q, sid)
			s.lock.Unlock()
		}(ctx, nc, sid)
	}
	log.Debug("check active connections...")
	s.lock.Lock()
	defer s.lock.Unlock()
	count := 0
	for sid, nc := range s.q {
		log.Debugf("%s close connection", sid)
		nc.Close()
		count++
	}
	s.q = nil
	s.q = make(map[string]net.Conn)
	log.Debugf("closed %d active connection(s)", count)
	return nil
}

func (s *Console) dispatch(ctx context.Context, nc net.Conn, sid string) {
	log.Debugf("%s dispatch", sid)
	select {
	case <-ctx.Done():
		log.Debugf("%s dispatch context done: %v", sid, ctx.Err())
		return
	default:
	}
	// ssh handshake
	conn, chans, reqs, err := ssh.NewServerConn(nc, s.cfg)
	if err != nil {
		log.Debugf("%s handshake error: %v", sid, err)
		return
	}
	defer conn.Close()
	s.wg.Add(1)
	go func(reqs <-chan *ssh.Request) {
		defer s.wg.Done()
		ssh.DiscardRequests(reqs)
	}(reqs)
	// serve
	fp := conn.Permissions.Extensions["pubkey-fp"]
	log.Printf("Auth login %s %s", fp, sid)
LOOP:
	for nc := range chans {
		select {
		case <-ctx.Done():
			log.Debugf("%s serve context done: %v", sid, ctx.Err())
			break LOOP
		default:
		}
		s.serve(ctx, nc, sid)
	}
	log.Printf("Auth logout %s %s", sid, s.ctxSessionElapsed(ctx))
}

func (s *Console) serve(ctx context.Context, nc ssh.NewChannel, sid string) {
	log.Debugf("serve session %s", sid)
	t := nc.ChannelType()
	log.Debugf("%s serve channel type %s", sid, t)
	if t != "session" {
		nc.Reject(ssh.UnknownChannelType, "unknown channel type")
		log.Errorf("Console %s unknown channel type: %s", sid, t)
		return
	}
	ch, reqs, err := nc.Accept()
	if err != nil {
		log.Errorf("Console %s could not accept channel: %v", sid, err)
		return
	}
	s.wg.Add(1)
	go func(ctx context.Context, in <-chan *ssh.Request) {
		log.Debugf("%s serve request", sid)
		defer s.wg.Done()
		for req := range in {
			select {
			case <-ctx.Done():
				log.Debug("serve request context done!")
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
	}(ctx, reqs)
	term := terminal.NewTerminal(ch, fmt.Sprintf("%s> ", env.Get("MUNBOT")))
	s.wg.Add(1)
	go func(ctx context.Context, ch ssh.Channel) {
		log.Debugf("%s serve shell", sid)
		defer s.wg.Done()
		defer ch.Close()
		for {
			select {
			case <-ctx.Done():
				log.Debug("serve shell context done!")
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
