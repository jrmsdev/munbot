// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package hash implements some hashing utils.
package hash

import (
	"crypto/sha256"
	"fmt"
)

func Sum(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
