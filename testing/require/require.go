// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package require wraps github.com/stretchr/testify/require package.
package require

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Assertions wraps require.Assertions.
type Assertions struct {
	*require.Assertions
}

// New creates a new require.Assertions wrapper.
func New(t *testing.T) *Assertions {
	return &Assertions{require.New(t)}
}
