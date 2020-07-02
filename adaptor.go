// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"github.com/jrmsdev/munbot/log"
)

type Adaptor struct {
	name string
}

func NewAdaptor() *Adaptor {
	return &Adaptor{
		name: "munbot",
	}
}

func (a *Adaptor) Name() string {
	return a.name
}

func (a *Adaptor) SetName(name string) {
	a.name = name
}

func (a *Adaptor) Connect() error {
	log.Print("Connect adaptor", a.name, "...")
	return nil
}

func (a *Adaptor) Finalize() error {
	log.Print("Finalize adaptor", a.name, "...")
	return nil
}

func (a *Adaptor) Ping() string {
	return "pong"
}
