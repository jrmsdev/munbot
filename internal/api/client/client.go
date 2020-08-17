// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package client defines api's client interfaces.
package client

import (
	"net/http"
)

// based on https://golangbyexample.com/command-design-pattern-in-golang/

type Receiver interface {
	GET(url string) (*http.Response, error)
	POST(url, content string) (*http.Response, error)
}

type Commander interface {
	Exec(*Request) Response
}

type Invoker interface {
	Eval(string) error
}
