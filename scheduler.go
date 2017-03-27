package main

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/michaelklishin/rabbit-hole"
	"github.com/maikelh/rabbitmq-metrics-exporter/exporters"
)

type SchedulerInterface interface {
	Start()
}

type Scheduler struct {
	ticker *time.Ticker
	rabbit *rabbithole.Client
	exporter exporters.Exporter
}

func (s *Scheduler) Start() error {
	s.rabbit, _ = rabbithole.NewClient("http://127.0.0.1:15672", "guest", "guest")
	export, err := exporters.NewStatsDExporter("localhost")

	if err != nil {
		return err
	}

	s.exporter = export

	s.ticker = time.NewTicker(time.Second * 10)
	for {
		select {
		case tickTime := <-s.ticker.C:
			s.tickHandler(tickTime)
		}
	}
}

func (s *Scheduler) tickHandler(time time.Time) {
	// For now we only handle queues, other info can come later
	queues, err := s.rabbit.ListQueues()

	if err != nil {
		logrus.Error(err)
		return
	}

	err = s.exporter.UpdateQueues(queues, "localhost", "/", time)

	if err != nil {
		logrus.Error(err)
	}
}