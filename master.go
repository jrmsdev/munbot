// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"gobot.io/x/gobot"

	"github.com/jrmsdev/munbot/internal/master/api"
)

type Master struct {
	*gobot.Master
	api *api.Api
}

func New() *Master {
	m := gobot.NewMaster()
	return &Master{m, api.New(m)}
}
