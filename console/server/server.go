// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package server implements master's console ssh server.
package server

// Server implements the ssh console server.
type Server struct {
}

func New() *Server {
	return &Server{}
}
