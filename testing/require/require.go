// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

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
