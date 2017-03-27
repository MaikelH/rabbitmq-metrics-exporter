package exporters

import (
	"github.com/michaelklishin/rabbit-hole"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/cactus/go-statsd-client/statsd"
	"strconv"
	"fmt"
)

type StatsDExporter struct {
	Host string
	Port int
	client statsd.Statter
}

func NewStatsDExporter(host string) (*StatsDExporter, error) {
	g := new(StatsDExporter)
	g.Host = host
	g.Port = 8125

	err := g.setupStatsD()

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *StatsDExporter) setupStatsD() error  {
	//stats, err := statsd.NewBufferedClient(g.Host + ":" + strconv.Itoa(g.Port), "rabbitmq", 300*time.Millisecond ,0)
	logrus.Info(fmt.Sprintf(g.Host + ":" + strconv.Itoa(g.Port)))

	stats, err := statsd.NewClient(g.Host + ":" + strconv.Itoa(g.Port), "rabbtimq")

	if err != nil {
		return err
	}

	g.client = stats

	return nil
}

func (g *StatsDExporter) UpdateQueues(queues []rabbithole.QueueInfo, host string, vhost string, time time.Time) error {
	var prefix = "queues."

	for _,queue := range queues {
		var queuePrefix = prefix + queue.Name

		g.client.Inc(queuePrefix + ".messages.total", int64(queue.Messages), 1.0)
		g.client.Inc(queuePrefix + ".messages.ready", int64(queue.MessagesReady), 1.0)
	}

	logrus.Info("Sending metrics to StatsD")

	return nil
}
