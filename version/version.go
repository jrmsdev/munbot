// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package version holds the running version information.
package version

import (
	"fmt"
)

const (
	Major int = 0
	Minor int = 0
	Patch int = 200710
)

func String() string {
	v := new(Info)
	b := v.Build().String()
	if b != "" {
		b = " (" + b + ")"
	}
	return fmt.Sprintf("%s%s", v, b)
}

var printf func(string, ...interface{}) (int, error)

func init() {
	printf = fmt.Printf
}

func Print(progname string) {
	printf("%s version %s\n", progname, String())
}
