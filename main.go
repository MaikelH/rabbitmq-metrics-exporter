package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting RabbitMQ Metrics Exporter")
	var c = new(Scheduler)

	c.Start()
}