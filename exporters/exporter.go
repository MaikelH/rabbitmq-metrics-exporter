package exporters

import (
	"github.com/michaelklishin/rabbit-hole"
	"time"
)

type Exporter interface {
	UpdateQueues(queues []rabbithole.QueueInfo, host string, vhost string, time time.Time) error
}
