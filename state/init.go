// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

var InitPanic = errors.New("init: run Master.Init first")
var InitError = errors.New("init: runtime already done")

type InitState struct {
	m   *sm
	err error
}

func newInit(m *sm) *InitState {
	return &InitState{m: m}
}

func (s *InitState) String() string {
	return "Init"
}

func (s *InitState) Error() error {
	return s.err
}

func (s *InitState) Run(ctx context.Context) Status {
	select {
	case <-ctx.Done():
		return DONE
	default:
		if s.m.Config == nil || s.m.ConfigFlags == nil || s.m.CoreFlags == nil {
			s.err = InitPanic
			log.Error(s.err)
			return PANIC
		}
		if s.m.Runtime != nil {
			s.err = InitError
			log.Error(s.err)
			return ERROR
		}
		s.m.Runtime = core.NewRuntime(ctx)
		if err := s.m.Runtime.Lock(); err != nil {
			s.err = err
			log.Error(s.err)
			return ERROR
		}
		log.Debug(s.m.Runtime)
		s.m.setState(s.m.configure)
	}
	return OK
}
