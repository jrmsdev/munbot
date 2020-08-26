// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package ssh implements internal ssh utilities.
package ssh

import (
	"context"
	"time"

	"github.com/munbot/master/v0/utils/net"
	"github.com/munbot/master/v0/utils/uuid"
)

type Server interface {
	Configure() error
	Start() error
	Stop() error
	Addr() *net.Addr
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
