// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
)

var (
	buildDate string = " "
	buildOS   string = " "
	buildArch string = " "
	buildUser string = " "
	buildHost string = " "
)

func buildInfo() string {
	if buildDate != " " {
		return fmt.Sprintf(" (%s)", &Build{})
	}
	return ""
}

type Build struct{}

func (b *Build) String() string {
	return fmt.Sprintf("%s %s/%s %s@%s",
		buildDate, buildOS, buildArch, buildUser, buildHost)
}
