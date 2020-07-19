// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package require is just a handy shorcut (?) to import
// github.com/stretchr/testify/require functionality.
package require

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Assertions struct {
	*require.Assertions
}

func New(t *testing.T) *Assertions {
	return &Assertions{require.New(t)}
}
