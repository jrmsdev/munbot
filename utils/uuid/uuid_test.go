// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package uuid

import (
	"testing"

	"github.com/munbot/master/testing/assert"
)

func TestLocker(t *testing.T) {
	assert := assert.New(t)
	assert.Regexp("^([[:xdigit:]]*)-([[:xdigit:]]*)-([[:xdigit:]]*)-([[:xdigit:]]*)-([[:xdigit:]]*)$", Rand())
}
