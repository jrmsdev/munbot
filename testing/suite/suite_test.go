// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/munbot/master/testing/require"
)

func TestSuite(t *testing.T) {
	require := require.New(t)
	require.Implements((*suite.TestingSuite)(nil), New())
}

type mockSuite struct {
	*suite.Suite
}

func newMockSuite() *mockSuite {
	return &mockSuite{Suite: new(suite.Suite)}
}

func TestRun(t *testing.T) {
	s := newMockSuite()
	Run(t, s)
}
