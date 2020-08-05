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
	x := fmt.Sprintf("%s %s/%s [%s]", buildDate, buildOS, buildArch, buildTags)
	assert.Equal(x, s)
}

func TestBuildArgs(t *testing.T) {
	assert := assert.New(t)
	b := new(Build)
	assert.Equal("nodate", b.Date())
	assert.Equal("noos", b.OS())
	assert.Equal("noarch", b.Arch())
	assert.Equal([]string{"notags"}, b.Tags())
	buildTags = "tag1,tag2,tag3"
	defer func() { buildTags = "notags" }()
	assert.Equal([]string{"tag1", "tag2", "tag3"}, b.Tags())
}
