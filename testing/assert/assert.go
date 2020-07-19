// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package assert is just a handy shorcut (?) to import
// github.com/stretchr/testify/assert functionality.
package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Assertions struct {
	*assert.Assertions
}

func New(t *testing.T) *Assertions {
	return &Assertions{assert.New(t)}
}
