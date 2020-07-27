// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// uuid utils.
package uuid

import (
	"github.com/gofrs/uuid"
)

func Rand() string {
	r := uuid.Must(uuid.NewV4())
	return r.String()
}
