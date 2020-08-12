// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package console

import (
	"context"

	"github.com/munbot/master/utils/uuid"
)

type ctxKey int

const (
	ctxSession ctxKey = iota
)

func (s *Console) ctxNewSession(ctx context.Context) (context.Context, string) {
	sid := uuid.Rand()
	return context.WithValue(ctx, ctxSession, sid), sid
}

func (s *Console) ctxSession(ctx context.Context) string {
	return ctx.Value(ctxSession).(string)
}
