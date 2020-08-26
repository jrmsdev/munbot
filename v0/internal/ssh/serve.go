// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/session"
	"github.com/munbot/master/v0/internal/user"
	"github.com/munbot/master/v0/log"
)

type request struct {
	Type  string
	Serve bool
}

func (s *SSHD) serve(ctx context.Context, nc ssh.NewChannel, sid session.Token, uid user.ID, fp string) {
	log.Debugf("session %s", sid)
	t := nc.ChannelType()
	log.Debugf("%s channel type %s", sid, t)
	if t != "session" {
		nc.Reject(ssh.UnknownChannelType, "unknown channel type")
		log.Errorf("SSHD %s unknown channel type: %s", sid, t)
		return
	}
	ch, chr, err := nc.Accept()
	if err != nil {
		log.Errorf("SSHD %s could not accept channel: %v", sid, err)
		return
	}
	reqs := make(chan request, 1)
	s.wgadd("serve-request")
	go func(ctx context.Context, in <-chan *ssh.Request, out chan<- request) {
		log.Debugf("%s serve request", sid)
		defer s.wgdone("serve-request")
		for req := range in {
			select {
			case <-ctx.Done():
				log.Debug("request context done!")
				req.Reply(false, nil)
				return
			default:
			}
			log.Debugf("%s request type %s", sid, req.Type)
			serve := false
			switch req.Type {
			case "pty-req", "env", "shell":
				serve = true
			}
			req.Reply(serve, nil)
			out <- request{Type: req.Type, Serve: serve}
		}
	}(ctx, chr, reqs)
	s.wgadd("serve")
	go func(ctx context.Context, ch ssh.Channel, in <-chan request) {
		defer s.wgdone("serve")
		wait := true
		for wait {
			req := <-in
			switch req.Type {
			case "shell":
				wait = false
				s.serveShell(ctx, ch, sid, uid, fp)
			default:
				if !req.Serve {
					log.Errorf("%s ssh invalid request: %s", sid, req.Type)
					wait = false
					if err := ch.Close(); err != nil {
						log.Errorf("%s ssh channel close: %v", sid, err)
					}
				}
			}
		}
	}(ctx, ch, reqs)
}

func (s *SSHD) serveShell(ctx context.Context, ch ssh.Channel, sid session.Token, uid user.ID, fp string) {
	log.Debugf("%s serve shell", sid)
	defer func() {
		if err := ch.Close(); err != nil {
			if err != io.EOF {
				log.Errorf("SSHD channel close %s: %v", sid, err)
			}
		}
	}()
	term := terminal.NewTerminal(ch, "")
	resp := bufio.NewWriter(term)
	if err := s.auth.Login(sid, uid, fp); err != nil {
		log.Debugf("%s auth login error: %v", sid, err)
		if err := shellWrite(resp, "", "login error"); err != nil {
			log.Errorf("SSHD terminal %s: %v", sid, err)
		}
		return
	}
	shellWrite(resp, "", "login")
	ps1 := fmt.Sprintf("%s> ", env.Get("MUNBOT"))
	term.SetPrompt(ps1)
LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Debugf("%s shell context done!", sid)
			break LOOP
		default:
		}
		log.Debugf("%s shell read line...", sid)
		line, err := term.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Errorf("SSHD terminal %s: %v", sid, err)
				return
			}
			break LOOP
		} else {
			log.Printf("%s SHELL: %s", sid, line)
			if err := shellWrite(resp, ps1, "%q", line); err != nil {
				log.Errorf("SSHD terminal %s: %v", sid, err)
				return
			}
		}
	}
	term.SetPrompt("")
	if err := shellWrite(resp, ps1, "logout"); err != nil {
		log.Errorf("SSHD terminal %s: %v", sid, err)
	}
}

func shellWrite(w *bufio.Writer, ps, f string, args ...interface{}) error {
	if _, err := w.WriteString(ps + fmt.Sprintf(f, args...) + "\n"); err != nil {
		if err != io.EOF {
			return err
		}
	}
	if err := w.Flush(); err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}
