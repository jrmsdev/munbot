// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package ssh implements internal ssh utilities.
package ssh

import (
	"context"
	"net"
	"net/url"
	"time"

	"github.com/munbot/master/v0/utils/uuid"
)

type Server interface {
	Configure() error
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

type ctxKey int

const (
	ctxSession ctxKey = iota
	ctxBorn
)

func (s *SSHD) ctxNewSession(ctx context.Context) (context.Context, string) {
	sid := uuid.Rand()
	ctx = context.WithValue(ctx, ctxBorn, time.Now())
	return context.WithValue(ctx, ctxSession, sid), sid
}

func (s *SSHD) ctxSession(ctx context.Context) string {
	return ctx.Value(ctxSession).(string)
}

func (s *SSHD) ctxSessionElapsed(ctx context.Context) time.Duration {
	born := ctx.Value(ctxBorn).(time.Time)
	return time.Since(born)
}
