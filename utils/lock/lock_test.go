// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lock

import (
	"testing"

	"github.com/subchen/go-trylock/v2"

	"github.com/munbot/master/testing/require"
)

func TestLocker(t *testing.T) {
	require := require.New(t)
	require.Implements((*trylock.TryLocker)(nil), New())
}
