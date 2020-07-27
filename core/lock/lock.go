// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Lock core runtime.
package lock

import (
	"context"
	"errors"
	"time"

	"github.com/munbot/master/utils/lock"
	"github.com/munbot/master/utils/uuid"
)

type key int

const lockKey key = 0
var lockUUID string = ""
var mu *lock.Locker = lock.New()

func NewContext(ctx context.Context) (context.Context, error) {
	var err error
	lockUUID, err = tryLock()
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, lockKey, lockUUID), nil
}

func tryLock() (string, error) {
	if mu.TryLockTimeout(time.Second) {
		return uuid.Rand(), nil
	}
	return "", errors.New("core lock timeout")
}
