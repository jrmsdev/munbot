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
	buildTags string = " "
)

func buildInfo() string {
	if buildDate != " " {
		return fmt.Sprintf(" (%s %s/%s)", buildDate, buildOS, buildArch)
	}
	return ""
}
