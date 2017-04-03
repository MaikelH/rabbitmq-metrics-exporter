package main

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/michaelklishin/rabbit-hole"
	"github.com/maikelh/rabbitmq-metrics-exporter/exporters"
	"fmt"
	"github.com/spf13/viper"
	"github.com/maikelh/rabbitmq-metrics-exporter/structs"
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
	var rabbitmqUrl = fmt.Sprintf("http://%s:%s", viper.GetString("rabbitmq.host"), viper.GetString("RabbitMQ.port"))

	s.rabbit, _ = rabbithole.NewClient(rabbitmqUrl, viper.GetString("RabbitMQ.user"), viper.GetString("RabbitMQ.password"))
	export, err := exporters.CreateExporter(viper.GetString("exporter.type"))

	if err != nil {
		return err
	}

	s.exporter = export

	s.ticker = time.NewTicker(time.Second * 5)
	for {
		select {
		case tickTime := <-s.ticker.C:
			s.tickHandler(tickTime)
		}
	}
}

func (s *Scheduler) tickHandler(time time.Time) {

	// For now we only handle queues, other info can come later
	queues, err := s.getQueueInformation()

	if err != nil {
		logrus.Error(err)
		return
	}

	err = s.exporter.UpdateQueues(queues, viper.GetString("rabbitmq.host"), "/", time)

	if err != nil {
		logrus.Error(err)
	}
}

func (s *Scheduler) getQueueInformation() ([]structs.Queue, error) {
	rabbitQueues, err := s.rabbit.ListQueues()

	if err != nil {
		return nil, err
	}

	var queues []structs.Queue

	for _, rabbitQueue := range rabbitQueues {
		var queue = structs.Queue{}

		queue.Name = rabbitQueue.Name
		queue.Node = rabbitQueue.Node
		queue.Vhost = rabbitQueue.Vhost
		queue.AutoDelete = rabbitQueue.AutoDelete
		queue.Durable = rabbitQueue.Durable

		queue.MessagesTotal = int64(rabbitQueue.Messages)
		queue.MessagesReady = int64(rabbitQueue.MessagesReady)
		queue.MessagesUnacknowledged = int64(rabbitQueue.MessagesUnacknowledged)
		queue.MessageBytes = int64(rabbitQueue.MessagesBytes)
		queue.MessageBytesReady = int64(rabbitQueue.MessagesBytes)
		queue.MessagesRAM = int64(rabbitQueue.MessagesRAM)
		queue.MessagesPersistent = int64(rabbitQueue.MessagesPersistent)

		queue.RateDelivered = int64(rabbitQueue.MessageStats.DeliverDetails.Rate)
		queue.RateDeliveredGet = int64(rabbitQueue.MessageStats.DeliverGetDetails.Rate)
		queue.RateDeliveredNoAck = int64(rabbitQueue.MessageStats.DeliverNoAckDetails.Rate)
		queue.RatePublished = int64(rabbitQueue.MessageStats.PublishDetails.Rate)
		queue.RateRedelivered = int64(rabbitQueue.MessageStats.RedeliverDetails.Rate)

	}

	return queues, nil
}

