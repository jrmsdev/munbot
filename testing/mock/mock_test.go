// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mock

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/munbot/master/testing/require"
)

func TestAssertions(t *testing.T) {
	check := require.New(t)
	check.IsType(new(mock.Mock), New())
}
