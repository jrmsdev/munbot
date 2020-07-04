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
	if Patch > 0 {
		return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	}
	return fmt.Sprintf("%d.%d", Major, Minor)
}

func Print(progname string) {
	fmt.Printf("%s version %s%s\n", progname, String(), buildInfo())
}
