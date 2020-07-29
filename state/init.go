// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

	"github.com/munbot/master/log"
)

var _ State = &InitState{}


type InitState struct {
	m   Machine
	id  StateID
	err error
}

func NewInitState(m Machine) *InitState {
	return &InitState{m: m, id: Init}
}

func (s *InitState) String() string {
	return s.id.String()
}

func (s *InitState) Error() error {
	return s.err
}

var InitPanic error = errors.New("init: run Master.Init first")

func (s *InitState) Run(ctx context.Context) (context.Context, Status) {
	select {
	case <-ctx.Done():
		return ctx, DONE
	default:
		if s.m.Config() == nil || s.m.ConfigFlags() == nil || s.m.CoreFlags() == nil {
			s.err = InitPanic
			log.Error(s.err)
			return ctx, PANIC
		}
		rt := s.m.Runtime()
		var err error
		if ctx, err = rt.Lock(ctx); err != nil {
			s.err = err
			log.Error(s.err)
			return ctx, ERROR
		}
		log.Debug(rt)
		s.m.SetState(Configure)
	}
	return ctx, OK
}
