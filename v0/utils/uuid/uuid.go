// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package uuid wraps github.com/gofrs/uuid package.
package uuid

import (
	"github.com/gofrs/uuid"
)

type UUID uuid.UUID

var Nil UUID = UUID{}

// Rand returns a new random UUID string.
func Rand() string {
	return uuid.Must(uuid.NewV4()).String()
}

func FromString(input string) (UUID, error) {
	var u UUID
	if b, err := uuid.FromString(input); err != nil {
		return Nil, err
	} else {
		u = UUID(b)
	}
	return u, nil
}

func ToString(input UUID) string {
	return uuid.UUID(input).String()
}
