// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()
	c.require.Equal("munbot", c.test.Get("master", "name"), "master name default")

	err := Update(c.test, "master.name", "testing")
	c.require.NoError(err, "update master.name error")
	c.assert.Equal("testing", c.test.Get("master", "name"), "master name update")
}
