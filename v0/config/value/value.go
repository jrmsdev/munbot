// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package value holds the maps for the config parser.
package value

// Map holds config section options with its values.
type Map map[string]string

// DB maps a config section to its values.
type DB map[string]Map
