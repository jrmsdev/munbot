// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
	"testing"

	"github.com/munbot/master/testing/assert"
)

func init() {
	buildDate = "nodate"
	buildOS = "noos"
	buildArch = "noarch"
	buildTags = "notags"
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	s := String()
	x := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	assert.Equal(x, s)
	buildDate = "date"
	defer func() { buildDate = "nodate" }()
	s = String()
	x = fmt.Sprintf("%s (%s %s/%s [%s])", x,
		buildDate, buildOS, buildArch, buildTags)
	assert.Equal(x, s)
}

func TestPrint(t *testing.T) {
	assert := assert.New(t)
	mockOut := "nooutput"
	mockPrintf := func(f string, a ...interface{}) (int, error) {
		mockOut = fmt.Sprintf(f, a...)
		return 0, nil
	}
	printf = mockPrintf
	defer func() { printf = fmt.Printf }()
	Print("testing")
	x := fmt.Sprintf("testing version %s\n", String())
	assert.Equal(x, mockOut)
}
