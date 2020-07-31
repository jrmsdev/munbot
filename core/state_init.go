// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"

	//~ "github.com/munbot/master/config"
)

var ErrSInit error = errors.New("core: run Init first")

var _ State = &SInit{}

type SInit struct {
	m Machine
	rt *Mem
}

func NewInit(m Machine, rt *Mem) State {
	return &SInit{m: m, rt: rt}
}

func (s *SInit) Init() error {
	return nil
}

func (s *SInit) Configure() error {
	return ErrSInit
}

func (s *SInit) Start() error {
	return ErrSInit
}

func (s *SInit) Stop() error {
	return ErrSInit
}
