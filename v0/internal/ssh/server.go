// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh

import (
	"context"
	"fmt"
	"net"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/munbot/master/v0/config/profile"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/session"
	"github.com/munbot/master/v0/internal/user"
	"github.com/munbot/master/v0/log"
)

var _ Server = &SSHD{}

// SSHD implements the ssh server.
type SSHD struct {
	enable bool
	auth   AuthManager
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

func NewServer() *SSHD {
	return &SSHD{
		done:   make(chan bool, 1),
		wg:     &sync.WaitGroup{},
		lock:   new(sync.Mutex),
		q:      make(map[string]net.Conn),
		closed: true,
		wgc:    make(map[string]int),
	}
}

func (s *SSHD) wgadd(n string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.wg.Add(1)
	s.wgc[n] += 1
}

func (s *SSHD) wgdone(n string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.wg.Done()
	s.wgc[n] -= 1
}

func (s *SSHD) wgwait() {
	log.Debugf("wgwait %v", s.wgc)
	s.wg.Wait()
}

func (s *SSHD) Addr() *Addr {
	addr := ""
	if s.ln != nil {
		addr = s.ln.Addr().String()
	}
	return newAddr("tcp", addr)
}

func (s *SSHD) Configure() error {
	s.enable = env.GetBool("MBSSHD")
	if s.enable {
		if s.auth == nil {
			s.auth = NewServerAuth()
		}
		p := profile.New()
		if err := s.auth.Configure(p.GetPath("auth")); err != nil {
			log.Debugf("auth manager configure error: %v", err)
			return err
		}
		s.cfg = s.auth.ServerConfig()
		s.addr = fmt.Sprintf("%s:%d", env.Get("MBSSHD_ADDR"), env.GetUint("MBSSHD_PORT"))
	}
	return nil
}

func (s *SSHD) Stop() error {
	log.Debug("stop")
	if s.enable && !s.closed {
		log.Print("Stop ssh server.")
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
	return nil
}

func (s *SSHD) Start() error {
	log.Debug("start")
	if s.enable {
		var err error
		// listen
		s.ln, err = net.Listen("tcp", s.addr)
		if err != nil {
			log.Debugf("listen error: %v", err)
			return err
		}
		log.Infof("SSH server ssh://%s", s.addr)
		s.closed = false
		// accept connections
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		defer s.ln.Close()
		return s.accept(ctx)
	} else {
		log.Warn("SSH server is disabled")
	}
	return nil
}

func (s *SSHD) accept(ctx context.Context) error {
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
				log.Errorf("SSH server accept: %v", err)
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

func (s *SSHD) dispatch(ctx context.Context, nc net.Conn, sid string) {
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
	uid := conn.Permissions.Extensions["x-munbot-user"]
	if fp == "" || uid == "" {
		log.Errorf("Auth invalid credentials %s.", sid)
		return
	}
	var sess session.Token
	if sess, err = session.FromString(sid); err != nil {
		log.Debugf("%s dispatch session error: %v", sid, err)
		return
	}
LOOP:
	for nc := range chans {
		select {
		case <-ctx.Done():
			log.Debugf("%s serve context done: %v", sid, ctx.Err())
			nc.Reject(ssh.ConnectionFailed, "context done")
			break LOOP
		default:
		}
		s.serve(ctx, nc, sess, user.ID(uid), fp)
	}
	if err := session.Close(sess); err != nil {
		log.Debugf("%s dispatch close session error: %v", sid, err)
	}
}
