package exporters

import (
	"time"
	"errors"
	"github.com/maikelh/rabbitmq-metrics-exporter/structs"
)

type Exporter interface {
	UpdateQueues(queues []structs.Queue, host string, vhost string, time time.Time) error
}

func CreateExporter(name string) (Exporter, error) {
	switch name {
	case "statsd":
		return NewStatsDExporter()
	case "console":
		return NewConsoleExporter()
	}


	return nil, errors.New("No valid exporter type given")
}