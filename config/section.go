// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"

	"github.com/munbot/master/config/internal/parser"
	"github.com/munbot/master/log"
)

// Section is the config section manager.
type Section struct {
	name string
	h    *parser.Config
}

// Name returns section's name.
func (s *Section) Name() string {
	return s.name
}

// HasOption checks if the named option exists in this section.
func (s *Section) HasOption(name string) bool {
	return s.h.HasOption(s.name, name)
}

// Get returns the evalualed (${var} expanded) content for the named option.
func (s *Section) Get(name string) string {
	return s.h.Get(s.name, name)
}

// GetBool returns the bool value for the named option.
// If option does not exists or if there's any parsing error, false will be
// returned as default value.
func (s *Section) GetBool(name string) bool {
	r, err := strconv.ParseBool(s.Get(name))
	if err != nil {
		log.Errorf("config option %s.%s parse error: %s", s.name, name, err)
		return false
	}
	return r
}

// GetInt returns the int value for the named option.
// Default value: 0.
func (s *Section) GetInt(name string) int {
	r, err := strconv.Atoi(s.Get(name))
	if err != nil {
		log.Errorf("config option %s.%s parse error: %s", s.name, name, err)
		return 0
	}
	return r
}
