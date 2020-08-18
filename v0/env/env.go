// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package env manages settings configurable from os.Environ and .env files.
//
// Gets config settings from os.Environ or from Defaults otherwise. It will also
// load .env files if present.
//
// Settings from .env file will be used to populate os.Environ at init() time.
//
// If MBENV_CONFIG is set as "" (empty) no more files will be loaded, which is
// the default behavior.
//
// But if file ${MBENV_CONFIG}/${MBENV}.env exists it will be loaded too.
//
// Every loaded file overrides the settings from previous one.
package env

import (
	"strconv"

	"github.com/gobuffalo/envy"

	"github.com/munbot/master/v0/log"
)

// Get key value, using Init copy for its default value. If not present, returns
// UNSET.
func Get(key string) string {
	return envy.Get(key, defvalGet(key))
}

// GetBool returns the bool value for key.
// If there's a parsing error it will be logged and return default value false.
func GetBool(key string) bool {
	r, err := strconv.ParseBool(Get(key))
	if err != nil {
		log.Errorf("env parse bool %s: %s", key, err)
		return false
	}
	return r
}

// GetInt returns the int value for key.
// If there's a parsing error it will be logged and return default value 0.
func GetInt(key string) int {
	r, err := strconv.Atoi(Get(key))
	if err != nil {
		log.Errorf("env parse int %s: %s", key, err)
		return 0
	}
	return r
}

// GetUint returns the uint value for key.
// If there's a parsing error it will be logged and return default value 0.
func GetUint(key string) uint {
	r, err := strconv.ParseUint(Get(key), 10, 0)
	if err != nil {
		log.Errorf("env parse uint %s: %s", key, err)
		return 0
	}
	return uint(r)
}

// SetDefault sets a default value. Env is not modified, the option is added to
// the default settings. If it already exists, its value is updated.
func SetDefault(key, val string) {
	defvalSet(key, val)
}

// Set sets env key value. But it does not modify os.Environ.
func Set(key, val string) {
	envy.Set(key, val)
}

// SetInt sets an int value.
func SetInt(key string, val int) {
	envy.Set(key, strconv.FormatInt(int64(val), 10))
}

// SetUint sets an uint value.
func SetUint(key string, val uint) {
	envy.Set(key, strconv.FormatUint(uint64(val), 10))
}
