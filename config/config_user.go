// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/munbot/internal/config"
)

type User struct {
	Enable *config.BoolValue     `json:"enable,omitempty"`
	Name   *config.StringValue   `json:"name,omitempty"`
}

func newUser(s *config.Section, name string) *User {
	return &User{
		Enable: s.NewBool(name+".enable", true),
		Name:   s.NewString(name+".name", name),
	}
}
