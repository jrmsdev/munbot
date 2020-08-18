// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"testing"

	"github.com/munbot/master/v0/testing/assert"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	v := Version()
	assert.Regexp(`^\d+\.\d+\.\d+$`, v.String())
}
