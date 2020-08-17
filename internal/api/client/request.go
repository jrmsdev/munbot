// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package client

// Response defines the api client's response interface.
type Request interface {
	Err() error
	FormatText() string
}
