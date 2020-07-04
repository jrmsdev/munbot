// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package version

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
