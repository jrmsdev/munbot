// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package assert wrapsgithub.com/stretchr/testify/assert package.
package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Assertions wraps assert.Assertions.
type Assertions struct {
	*assert.Assertions
}

// New creates a new assert.Assertions wrapper.
func New(t *testing.T) *Assertions {
	return &Assertions{assert.New(t)}
}
