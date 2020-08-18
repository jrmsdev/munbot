// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package suite wraps github.com/stretchr/testify/suite package.
package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Suite wraps suite.Suite.
type Suite struct {
	*suite.Suite
}

// New creates a new suite.Suite wrapper.
func New() *Suite {
	return &Suite{new(suite.Suite)}
}

// Run runs the suite.
func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}
