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
	fmt.Printf("%s version %s\n", progname, String())
}

type Info struct{}

func (v *Info) String() string {
	return String()
}

func (v *Info) Major() int {
	return Major
}

func (v *Info) Minor() int {
	return Minor
}

func (v *Info) Patch() int {
	return Patch
}
