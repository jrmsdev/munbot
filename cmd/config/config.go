// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
	"fmt"
	"sort"

	"github.com/munbot/master/cmd"
	"github.com/munbot/master/config"
	"github.com/munbot/master/config/profile"
	"github.com/munbot/master/log"
)

type Flags struct {
	ListAll bool
}

func (f *Flags) set(fs *flag.FlagSet) {
	fs.BoolVar(&f.ListAll, "a", false, "list all options")
}

type Cmd struct {
	flags *Flags
}

func New() *Cmd {
	return &Cmd{flags: &Flags{}}
}

func (c *Cmd) FlagSet(fs *flag.FlagSet) {
	c.flags.set(fs)
}

func (c *Cmd) Command(flags *config.Flags) cmd.Command {
	return &Main{flags: c.flags, profile: flags.Profile}
}

type Main struct {
	flags *Flags
	profile *profile.Profile
}

func (m *Main) Run(args []string) int {
	filter := ""
	if len(args) == 1 {
		filter = args[0]
	}
	return m.list(filter)
}

func (m *Main) list(filter string) int {
	cfg := config.New()
	if m.flags.ListAll || filter != "" {
		cfg.SetDefaults(config.Defaults)
	}
	if err := cfg.Load(m.profile); err != nil {
		log.Error(err)
		return 1
	}
	p := config.NewParser(cfg)
	pm := p.Map(filter)
	if v, ok := pm[filter]; ok {
		fmt.Printf("%s\n", v)
	} else {
		for _, k := range m.sort(pm) {
			fmt.Printf("%s=%s\n", k, pm[k])
		}
	}
	return 0
}

func (m *Main) sort(n map[string]string) []string {
	l := make([]string, 0, len(n))
	for k := range n {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}
