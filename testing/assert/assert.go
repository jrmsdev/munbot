// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func New(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
