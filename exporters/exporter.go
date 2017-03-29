package exporters

import (
	"github.com/michaelklishin/rabbit-hole"
	"time"
	"errors"
)

type Exporter interface {
	UpdateQueues(queues []rabbithole.QueueInfo, host string, vhost string, time time.Time) error
}

func CreateExporter(name string) (Exporter, error) {
	switch name {
	case "statsd":
		return NewStatsDExporter()
	}

	return nil, errors.New("No valid exporter type given")
}