// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Core runtime.
package core

var rt *runtime

func init() {
	rt = newRuntime()
}

type runtime struct {
}

func newRuntime() *runtime {
	return &runtime{}
}
