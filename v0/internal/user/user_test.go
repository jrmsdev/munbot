// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package user

import (
	"testing"

	"github.com/munbot/master/v0/testing/assert"
)

var tuser *jsonUser = &jsonUser{
	EName: "test",
	EAddr: "test@munbot.local",
	FP:    "12345678",
	ID:    "c5eaae084cea9fd6914ac08b7cbfe57a452ae8a0f11983e57975bacbbee4da66",
}

func TestMarshal(t *testing.T) {
	check := assert.New(t)
	s, err := Marshal("test@munbot.local", "12345678")
	check.NoError(err)
	check.Equal(tuser.String(), s)
}
