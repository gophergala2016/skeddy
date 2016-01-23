package main

import (
	"github.com/robfig/cron"
	"log"
)

type Scheduler struct {
	Instance *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{Instance: cron.New()}
}

func (s *Scheduler) Start() {
	log.Println("Starting scheduler ...")
	s.Instance.Start()
}

var Skeddy *Scheduler

func (s *Scheduler) AddEntry(exp, ep, p string) {
  log.Println("Adding entry", exp, ep, p)
  s.Instance.AddFunc(exp, func() { Dispatch(ep, p) })
}
