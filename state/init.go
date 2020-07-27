// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"errors"

	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

var InitError = errors.New("init: run Master.Configure first")
var InitDoneError = errors.New("init: runtime already done")

type InitState struct {
	m   *Machine
	err error
}

func newInit(m *Machine) *InitState {
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
		if s.m.Config == nil {
			s.err = InitError
			return PANIC
		}
		if s.m.Runtime != nil {
			s.err = InitDoneError
			return ERROR
		}
		s.m.Runtime = core.NewRuntime(ctx)
		if err := s.m.Runtime.Lock(); err != nil {
			s.err = err
			return ERROR
		}
		log.Debug(s.m.Runtime)
		s.m.setState(s.m.configure)
	}
	return OK
}
