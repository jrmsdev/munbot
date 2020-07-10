// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

import (
	"fmt"
)

type Info struct{}

func (v *Info) String() string {
	s := ""
	if Patch > 0 {
		s = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	} else {
		s = fmt.Sprintf("%d.%d", Major, Minor)
	}
	return s
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

func (v *Info) Build() *Build {
	return new(Build)
}
