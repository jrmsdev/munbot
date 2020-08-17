// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/munbot/master/env"
	"github.com/munbot/master/log"
)

type request struct {
	Type    string
	Serve   bool
}

func (s *Console) serve(ctx context.Context, nc ssh.NewChannel, sid string) {
	log.Debugf("session %s", sid)
	t := nc.ChannelType()
	log.Debugf("%s channel type %s", sid, t)
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
			case "pty-req", "shell":
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
				s.serveShell(ctx, ch, sid)
			default:
				if !req.Serve {
					log.Errorf("%s ssh invalid request: %s", sid, req.Type)
					wait = false
				}
			}
		}
	}(ctx, ch, reqs)
}

func (s *Console) serveShell(ctx context.Context, ch ssh.Channel, sid string) {
	log.Debugf("%s serve shell", sid)
	defer ch.Close()
	term := terminal.NewTerminal(ch, fmt.Sprintf("%s> ", env.Get("MUNBOT")))
	for {
		select {
		case <-ctx.Done():
			log.Debug("shell context done!")
			return
		default:
		}
		log.Debugf("%s shell read line...", sid)
		line, err := term.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Errorf("Console %s: %v", sid, err)
			}
			return
		}
		log.Printf("%s SHELL: %s", sid, line)
	}
}
