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

type Build struct {
	Date string
	OS   string
	Arch string
	Tags []string
}

func BuildInfo() *Build {
	return &Build{buildDate, buildOS, buildArch, strings.Split(buildTags, ",")}
}

func (b *Build) String() string {
	if b.Date != "nodate" {
		return fmt.Sprintf(" (%s %s/%s %v)", b.Date, b.OS, b.Arch, b.Tags)
	}
	return ""
}
