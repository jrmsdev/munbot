// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package core implements the runtime utils.
package core

import (
	"errors"

	"github.com/munbot/master/v0/utils/lock"
)

var LockPanic error = errors.New("core: lock failed")
var UnlockPanic error = errors.New("core: unlock failed")

var k *lock.Locker

func init() {
	k = lock.New()
}

func Lock() {
	if !k.TryLock(nil) {
		panic(LockPanic)
	}
}

func Unlock() {
	defer func() {
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if ok && msg == "Unlock() failed" {
				panic(UnlockPanic)
			} else {
				panic(r)
			}
		}
	}()
	k.Unlock()
}
