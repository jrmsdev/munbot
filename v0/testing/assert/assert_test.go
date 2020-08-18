// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/munbot/master/v0/testing/require"
)

func TestAssertions(t *testing.T) {
	check := require.New(t)
	check.IsType(new(assert.Assertions), New(t))
}
