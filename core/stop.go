// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

func (rt *Core) Stop() error {
	select {
	case <-rt.ctx.Done():
		return rt.ctx.Err()
	}
	return rt.state.Stop()
}