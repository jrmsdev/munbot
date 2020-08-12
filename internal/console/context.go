// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console

import (
	"context"
	"time"

	"github.com/munbot/master/utils/uuid"
)

type ctxKey int

const (
	ctxSession ctxKey = iota
	ctxBorn
)

func (s *Console) ctxNewSession(ctx context.Context) (context.Context, string) {
	sid := uuid.Rand()
	ctx = context.WithValue(ctx, ctxBorn, time.Now())
	return context.WithValue(ctx, ctxSession, sid), sid
}

func (s *Console) ctxSession(ctx context.Context) string {
	return ctx.Value(ctxSession).(string)
}

func (s *Console) ctxSessionElapsed(ctx context.Context) time.Duration {
	born := ctx.Value(ctxBorn).(time.Time)
	return time.Since(born)
}
