// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"sync"
	"time"

	"github.com/munbot/master/log"
)

var _ State = &SRun{}

type SRun struct {
	m  Machine
	rt *Mem
	wg *sync.WaitGroup
	start chan bool
	fail chan error
	wait time.Duration
}

func newRun(m Machine, rt *Mem) State {
	return &SRun{
		m: m,
		rt: rt,
		wg: new(sync.WaitGroup),
		start: make(chan bool, 1),
		fail: make(chan error),
		wait: 100*time.Millisecond,
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
	go func(wg *sync.WaitGroup, start chan bool, fail chan error) {
		defer wg.Done()
		<-start
		if err := s.rt.Master.Start(); err != nil {
			log.Debugf("start master robot: %s", err)
			fail <- err
		}
	}(s.wg, s.start, s.fail)
	time.Sleep(s.wait)
	select {
	case err := <-s.fail:
		return err
	default:
	}
	return nil
}

func (s *SRun) Run() error {
	var err error
	log.Print("Run")
	s.start <- true
	check := true
	for check {
		select {
		case err = <-s.fail:
			check = false
		default:
			time.Sleep(s.wait)
			if !s.rt.Master.Running() {
				log.Debug("master is not running... abort!")
				check = false
				select {
				case err = <-s.fail:
				default:
				}
			}
		}
	}
	s.wg.Wait()
	return err
}

func (s *SRun) Stop() error {
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
