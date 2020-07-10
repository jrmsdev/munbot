// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package require

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func New(t *testing.T) *require.Assertions {
	return require.New(t)
}
