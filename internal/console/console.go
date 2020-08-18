// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console server.
package console

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/config/profile"
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
	net.Addr
	addr string
	net  string
	uri  *url.URL
}

func newAddr(net, addr string) *Addr {
	return &Addr{addr: addr, net: net, uri: &url.URL{Scheme: "ssh", Host: addr}}
}

func (a *Addr) String() string {
	return a.addr
}

func (a *Addr) Network() string {
	return a.net
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
	wgc    map[string]int
}

func New() *Console {
	return &Console{
		done: make(chan bool, 1),
		wg:   &sync.WaitGroup{},
		lock: new(sync.Mutex),
		q:    make(map[string]net.Conn),
		wgc:  make(map[string]int),
	}
}

func (s *Console) wgadd(n string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.wg.Add(1)
	s.wgc[n] += 1
}

func (s *Console) wgdone(n string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.wg.Done()
	s.wgc[n] -= 1
}

func (s *Console) wgwait() {
	s.lock.Lock()
	log.Debugf("wgwait %v", s.wgc)
	s.lock.Unlock()
	s.wg.Wait()
}

func (s *Console) Addr() *Addr {
	addr := ""
	if s.ln != nil {
		addr = s.ln.Addr().String()
	}
	return newAddr("tcp", addr)
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
		log.Debug("wait for them to finish...")
		s.wgwait()
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
		defer s.ln.Close()
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
		s.wgadd("dispatch")
		go func(ctx context.Context, nc net.Conn, sid string) {
			defer s.wgdone("dispatch")
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
	s.wgadd("discard-requests")
	go func(reqs <-chan *ssh.Request) {
		defer s.wgdone("discard-requests")
		ssh.DiscardRequests(reqs)
	}(reqs)
	// serve
	fp := conn.Permissions.Extensions["pubkey-fp"]
	log.Infof("Auth login %s %s", fp, sid)
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
	log.Infof("Auth logout %s %s", sid, s.ctxSessionElapsed(ctx))
}
