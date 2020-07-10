// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
	"strings"
)

var (
	buildDate string = "nodate"
	buildOS   string = "noos"
	buildArch string = "noarch"
	buildTags string = "notags"
)

type Build struct{}

func (b *Build) String() string {
	if buildDate != "nodate" {
		return fmt.Sprintf("%s %s/%s [%s]", buildDate, buildOS, buildArch, buildTags)
	}
	return ""
}

func (b *Build) Date() string {
	return buildDate
}

func (b *Build) OS() string {
	return buildOS
}

func (b *Build) Arch() string {
	return buildArch
}

func (b *Build) Tags() []string {
	return strings.Split(buildTags, ",")
}
