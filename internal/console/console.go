// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/url"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/munbot/master/config/profile"
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

type Addr struct {
	uri *url.URL
}

func (a *Addr) String() string {
	return a.uri.String()
}

func (a *Addr) Hostname() string {
	return a.uri.Hostname()
}

func (a *Addr) Port() string {
	return a.uri.Port()
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

func (s *Console) Addr() *Addr {
	if s.ln != nil {
		uri, _ := url.Parse(fmt.Sprintf("ssh://%s", s.ln.Addr().String()))
		return &Addr{uri: uri}
	}
	return &Addr{uri: &url.URL{Scheme: "ssh"}}
}

func (s *Console) Configure(cfg *Config) error {
	if cfg.Enable {
		s.enable = true
		s.auth = cfg.Auth
		if s.auth == nil {
			p := profile.New()
			s.auth = auth.New()
			if err := s.auth.Configure(p.GetPath("auth")); err != nil {
				log.Debugf("auth manager configure error: %v", err)
				return err
			}
		}
		s.cfg = s.auth.ServerConfig()
		s.addr = fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
	}
	return nil
}

func (s *Console) Stop() error {
	log.Debug("stop")
	if s.enable {
		defer close(s.done)
		s.closed = true
		s.done <- true
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
	} else {
		log.Warn("Console server is disabled")
	}
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
		log.Debugf("new session %s", sid)
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
		count++
		log.Debugf("%s close connection #%d", sid, count)
		nc.Close()
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
			nc.Reject(ssh.ConnectionFailed, "context done")
			break LOOP
		default:
		}
		s.serve(ctx, nc, sid)
	}
	log.Printf("Auth logout %s %s", sid, s.ctxSessionElapsed(ctx))
}

type request struct {
	Type string
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
	ch, chr, err := nc.Accept()
	if err != nil {
		log.Errorf("Console %s could not accept channel: %v", sid, err)
		return
	}
	reqs := make(chan request, 0)
	s.wg.Add(1)
	go func(ctx context.Context, in <-chan *ssh.Request, out chan<- request) {
		log.Debugf("%s serve request", sid)
		defer s.wg.Done()
		for req := range in {
			select {
			case <-ctx.Done():
				log.Debug("serve request context done!")
				req.Reply(false, nil)
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
			out <- request{Type: req.Type}
		}
	}(ctx, chr, reqs)
	s.wg.Add(1)
	go func(ctx context.Context, ch ssh.Channel, in <-chan request) {
		defer s.wg.Done()
		defer ch.Close()
		req := <-in
		switch req.Type {
		case "pty-req", "shell":
			s.serveTerminal(ctx, ch, sid)
		default:
			log.Errorf("%s ssh invalid request: %s", sid, req.Type)
		}
	}(ctx, ch, reqs)
}

func (s *Console) serveTerminal(ctx context.Context, ch ssh.Channel, sid string) {
	term := terminal.NewTerminal(ch, fmt.Sprintf("%s> ", env.Get("MUNBOT")))
	log.Debugf("%s serve shell", sid)
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
}
