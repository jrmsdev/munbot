// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
	"testing"

	"github.com/munbot/master/testing/assert"
)

func TestBuildString(t *testing.T) {
	assert := assert.New(t)
	b := new(Build)
	buildDate = "date"
	defer func() { buildDate = "nodate" }()
	s := b.String()
	x := fmt.Sprintf("%s %s/%s", buildDate, buildOS, buildArch)
	assert.Equal(x, s)
}

func TestBuildArgs(t *testing.T) {
	assert := assert.New(t)
	b := new(Build)
	assert.Equal("nodate", b.Date())
	assert.Equal("noos", b.OS())
	assert.Equal("noarch", b.Arch())
	assert.Equal([]string{}, b.Tags())
	buildTags = "munbot,static"
	defer func() { buildTags = "notags" }()
	assert.Equal([]string{"static"}, b.Tags())
	buildTags = "munbot"
	assert.Equal([]string{}, b.Tags())
}
