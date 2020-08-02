// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package wapp implements the master api webapp.
package wapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type wapp struct {
	mux *mux.Router
}

func New() Api {
	r := mux.NewRouter()
	return &wapp{mux: r}
}

func (a *wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
