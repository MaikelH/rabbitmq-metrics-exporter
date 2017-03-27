package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting RabbitMQ Metrics Exporter")
	var c = new(Scheduler)

	err := c.Start()

	if err != nil {
		logrus.Error(err)
	}
}