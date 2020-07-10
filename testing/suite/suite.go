// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite suite.Suite

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}
