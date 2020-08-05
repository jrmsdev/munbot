// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
	"testing"

	"github.com/munbot/master/testing/assert"
)

func TestInfoString(t *testing.T) {
	assert := assert.New(t)
	v := new(Info)
	s := v.String()
	x := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	assert.Equal(x, s)
}

func TestInfoArgs(t *testing.T) {
	assert := assert.New(t)
	v := new(Info)
	assert.Equal(Major, v.Major())
	assert.Equal(Minor, v.Minor())
	assert.Equal(Patch, v.Patch())
	assert.IsType(new(Build), v.Build())
}
