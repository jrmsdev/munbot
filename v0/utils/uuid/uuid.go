// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package uuid wraps github.com/gofrs/uuid package.
package uuid

import (
	"github.com/gofrs/uuid"
)

// Rand returns a new random UUID string.
func Rand() string {
	return uuid.Must(uuid.NewV4()).String()
}
