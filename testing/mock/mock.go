// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mock wraps github.com/stretchr/testify/mock package.
package mock

import (
	"github.com/stretchr/testify/mock"
)

// Mock wraps mock.Mock.
type Mock struct {
	*mock.Mock
}

// New creates a new mock.Mock wrapper.
func New() *Mock {
	return &Mock{new(mock.Mock)}
}
