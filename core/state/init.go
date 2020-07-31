// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

//~ import (
	//~ "github.com/munbot/master/config"
//~ )

var _ State = &SInit{}

type SInit struct {
}

func (s *SInit) Init() error {
	return nil
}

func (s *SInit) Configure() error {
	return nil
}

func (s *SInit) Start() error {
	return nil
}

func (s *SInit) Stop() error {
	return nil
}
