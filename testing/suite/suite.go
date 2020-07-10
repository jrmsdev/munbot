// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	*suite.Suite
}

func New() *Suite {
	return &Suite{new(suite.Suite)}
}

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}
