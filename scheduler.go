package main

import (
	"github.com/robfig/cron"
	"log"
)

type Scheduler struct {
	Instance *cron.Cron
}

var Skeddy *Scheduler

func NewScheduler() *Scheduler {
	return &Scheduler{Instance: cron.New()}
}

func (s *Scheduler) Start(entries []*Entry) {
	log.Println("Starting scheduler ...")
	s.Instance.Start()
	for _, e := range entries {
		s.AddEntry(e)
	}
}

func (s *Scheduler) AddEntry(e *Entry) {
  log.Println("Adding entry", e)
  s.Instance.AddFunc(e.Expression, func() { Dispatch(e.Endpoint, e.Payload) })
}

func (s *Scheduler) ReStart(entries []*Entry) {
	log.Println("Restarting scheduler ...")
	s.Instance.Stop()
	s.Instance = cron.New()
	s.Instance.Start()
	for _, e := range entries {
		s.AddEntry(e)
	}
}

func StartScheduler() {
	Skeddy = NewScheduler()
	entries := Store.AllEntries()
	Skeddy.Start(entries)
}
