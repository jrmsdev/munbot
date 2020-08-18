// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package require

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssertions(t *testing.T) {
	check := New(t)
	check.IsType(new(require.Assertions), New(t))
}
