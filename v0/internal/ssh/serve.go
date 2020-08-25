// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ssh

import (
	//~ "bufio"
	"context"
	//~ "fmt"
	//~ "io"
	//~ "net/textproto"

	"golang.org/x/crypto/ssh"
	//~ "golang.org/x/crypto/ssh/terminal"

	//~ "github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/internal/session"
	"github.com/munbot/master/v0/log"
)

type request struct {
	Type  string
	Serve bool
}

func (s *SSHD) serve(ctx context.Context, nc ssh.NewChannel, sid session.Token) {
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
				if err := ch.Close(); err != nil {
					log.Errorf("%s ssh channel close: %v", sid, err)
				}
				//~ s.serveShell(ctx, ch, sid)
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

//~ func (s *SSHD) serveShell(ctx context.Context, ch ssh.Channel, sid string) {
	//~ log.Debugf("%s serve shell", sid)
	//~ defer ch.Close()
	//~ _, err := s.auth.Session(sid)
	//~ if err != nil {
		//~ log.Debugf("%s auth session error: %v", sid, err)
		//~ return
	//~ }
	//~ ps1 := fmt.Sprintf("%s> ", env.Get("MUNBOT"))
	//~ term := terminal.NewTerminal(ch, ps1)
	//~ resp := textproto.NewWriter(bufio.NewWriter(term))
//~ LOOP:
	//~ for {
		//~ select {
		//~ case <-ctx.Done():
			//~ log.Debug("shell context done!")
			//~ break LOOP
		//~ default:
		//~ }
		//~ log.Debugf("%s shell read line...", sid)
		//~ line, err := term.ReadLine()
		//~ if err != nil {
			//~ if err != io.EOF {
				//~ log.Errorf("SSHD %s: %v", sid, err)
				//~ return
			//~ }
			//~ break LOOP
		//~ }
		//~ log.Printf("%s SHELL: %s", sid, line)
		//~ if err := resp.PrintfLine("%q", line); err != nil {
			//~ log.Errorf("SSHD %s: %v", sid, err)
			//~ return
		//~ }
	//~ }
	//~ term.SetPrompt("")
	//~ if err := resp.PrintfLine("%s%s", ps1, "logout"); err != nil {
		//~ if err != io.EOF {
			//~ log.Errorf("SSHD %s: %v", sid, err)
		//~ }
	//~ }
//~ }
