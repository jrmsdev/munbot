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
	m     Machine
	rt    *Mem
	wg    *sync.WaitGroup
	start chan bool
	fail  chan failmsg
	wait  time.Duration
}

func newRun(m Machine, rt *Mem) State {
	return &SRun{
		m:     m,
		rt:    rt,
		wg:    new(sync.WaitGroup),
		start: make(chan bool),
		fail:  make(chan failmsg),
		wait:  300 * time.Millisecond,
	}
}

func (s *SRun) Init() error {
	return ErrInit
}

func (s *SRun) Configure() error {
	return ErrConfigure
}

func (s *SRun) Start() error {
	log.Print("Start")
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, start chan bool, fail chan failmsg) {
		defer wg.Done()
		<-start
		if err := s.rt.Master.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"master", err}
		}
	}(s.wg, s.start, s.fail)
	return nil
}

func (s *SRun) Run() error {
	log.Print("Run")
	var fail failmsg
	s.start <- true
	close(s.start)
LOOP:
	for {
		select {
		case fail = <-s.fail:
			break LOOP
		default:
			time.Sleep(s.wait)
			if !s.rt.Master.Running() {
				log.Debug("master is not running... abort!")
				select {
				case fail = <-s.fail:
				default:
					if err := s.rt.Api.Stop(); err != nil {
						log.Error(err)
					}
				}
				break LOOP
			}
		}
	}
	if fail.err != nil {
		log.Debugf("core %s failed: %s", fail.name, fail.err)
		if err := s.rt.Api.Stop(); err != nil {
			log.Errorf("stop master api: %s", err)
		}
		if err := s.rt.Master.Stop(); err != nil {
			log.Errorf("stop master robot: %s", err)
		}
	}
	s.wg.Wait()
	return fail.err
}

func (s *SRun) Stop() error {
	log.Print("Stop")
	var err error
	if err := s.rt.Master.Stop(); err != nil {
		err = log.Error(err)
	}
	s.wg.Wait()
	if err != nil {
		return err
	}
	return s.m.SetState(Halt)
}

func (s *SRun) Halt() error {
	return ErrHalt
}
