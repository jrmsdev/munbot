// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()

	s := c.test.Section("master")
	c.require.Equal(s.Get("name"), "munbot", "master name default")

	err := Update(c.test, "master.name", "testing")
	c.require.NoError(err, "update master.name error")

	s = c.test.Section("master")
	c.assert.Equal(s.Get("name"), "testing", "master name default")
}
