// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
)

const (
	Major int = 0
	Minor int = 0
	Patch int = 0
)

func String() string {
	v := ""
	if Patch > 0 {
		v = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	} else {
		v = fmt.Sprintf("%d.%d", Major, Minor)
	}
	return fmt.Sprintf("%s%s", v, BuildInfo())
}

func Print(progname string) {
	fmt.Printf("%s version %s\n", progname, String())
}
