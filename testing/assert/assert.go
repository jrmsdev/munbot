// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

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
