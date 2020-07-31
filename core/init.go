// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

func (rt *Core) Init() error {
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	}
	return nil
}
