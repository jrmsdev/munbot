// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"testing"
)

func TestStatusMap(t *testing.T) {
	if len(stMap) != int(lastStatus) {
		t.Errorf("len stMap(%d) != lastStatus(%d)", len(stMap), lastStatus)
	}
}
