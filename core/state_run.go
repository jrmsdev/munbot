// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"sync"
	"time"

	"github.com/munbot/master/log"
)

type failmsg struct {
	name string
	err  error
}

var _ State = &SRun{}

type SRun struct {
	m    Machine
	rt   *Mem
	wg   *sync.WaitGroup
	fail chan failmsg
	wait time.Duration
	exit chan bool
}

func newRun(m Machine, rt *Mem) State {
	return &SRun{
		m:    m,
		rt:   rt,
		wg:   new(sync.WaitGroup),
		fail: make(chan failmsg, 1),
		wait: 300 * time.Millisecond,
		exit: make(chan bool, 1),
	}
}

func (s *SRun) Init() error {
	return ErrInit
}

func (s *SRun) Configure() error {
	return ErrConfigure
}

func (s *SRun) Start() error {
	log.Print("Start...")
	// start master robot
	s.rt.Master.ExitNotify(s.exit)
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, fail chan failmsg) {
		defer wg.Done()
		log.Debug("robot start...")
		if err := s.rt.Master.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"robot", err}
		}
	}(s.wg, s.fail)
	time.Sleep(s.wait)
	// start api server
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, fail chan failmsg) {
		defer wg.Done()
		log.Debug("api start...")
		if err := s.rt.Api.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"api", err}
		}
	}(s.wg, s.fail)
	return nil
}

func (s *SRun) Run() error {
	log.Print("Run...")
	var fail failmsg
	abort := false
LOOP:
	for {
		select {
		case fail = <-s.fail:
			break LOOP
		case <-s.exit:
			log.Debug("master exit...")
			abort = true
			break LOOP
		default:
			time.Sleep(s.wait)
			if !s.rt.Master.Running() {
				log.Debug("master is not running...")
				select {
				case fail = <-s.fail:
					log.Debug("there was a failure before abort")
				default:
				}
				log.Debug("abort!")
				abort = true
				break LOOP
			}
		}
	}
	// close fail channel just in case...
	// maybe not a good idea? but at least it makes them panic...
	defer close(s.fail)
	if fail.err != nil {
		log.Debugf("core %s failed: %s", fail.name, fail.err)
		abort = true
	}
	if abort {
		return s.m.Abort()
	}
	return fail.err
}

func (s *SRun) Stop() error {
	log.Print("Stop...")
	// stop api
	log.Debug("try to stop api...")
	if err := s.rt.Api.Stop(); err != nil {
		log.Error(err)
	}
	// stop robot
	var err error
	if s.rt.Master.Running() {
		log.Debug("try to stop robot...")
		if err := s.rt.Master.Stop(); err != nil {
			err = log.Error(err)
		}
	}
	// wait for them...
	log.Debug("wait for them to finish...")
	s.wg.Wait()
	if err != nil {
		return err
	}
	return s.m.SetState(Halt)
}

func (s *SRun) Halt() error {
	return ErrHalt
}
