// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lock

import (
	"testing"

	"github.com/subchen/go-trylock/v2"

	"github.com/munbot/master/v0/testing/require"
)

func TestLocker(t *testing.T) {
	check := require.New(t)
	check.Implements((*trylock.TryLocker)(nil), New())
}
