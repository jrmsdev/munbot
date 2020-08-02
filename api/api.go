// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api defines and implements master api server interface.
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
	"github.com/munbot/master/log"
)

var serverTimeout time.Duration = 15 * time.Second
var stopTimeout time.Duration = 30 * time.Second

func newHTTPServer() *http.Server {
	return &http.Server{
		WriteTimeout: serverTimeout,
		ReadTimeout: serverTimeout,
	}
}

var _ Server = &Api{}

type Api struct {
	enable bool
	server *http.Server
}

func New() Server {
	return &Api{server: newHTTPServer()}
}

func (a *Api) Configure(kfl *flags.Flags, cfg *config.Section) error {
	a.enable = cfg.GetBool("enable")
	a.server.Addr = fmt.Sprintf("%s:%d", kfl.ApiAddr, kfl.ApiPort)
	return nil
}

func (a *Api) Start() error {
	log.Printf("Api server http://%s/", a.server.Addr)
	if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (a *Api) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != http.ErrServerClosed {
		return err
	}
	return nil
}
