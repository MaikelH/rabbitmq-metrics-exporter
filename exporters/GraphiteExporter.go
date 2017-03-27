package exporters

import (
	"github.com/michaelklishin/rabbit-hole"
	"github.com/sirupsen/logrus"
	"time"
	"strconv"
	"fmt"
	"github.com/marpaia/graphite-golang"
)

type GraphiteExporter struct {
	Host string
	Port int
	client *graphite.Graphite
}

func NewGraphiteExporter(host string) (*GraphiteExporter, error) {
	g := new(GraphiteExporter)
	g.Host = host
	g.Port = 2003

	err := g.setupGraphite()

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GraphiteExporter) setupGraphite() error  {
	Graphite, err := graphite.NewGraphiteWithMetricPrefix(g.Host, g.Port, "rabbitmq")

	if err != nil {
		return err
	}

	g.client = Graphite

	return nil
}

func (g *GraphiteExporter) UpdateQueues(queues []rabbithole.QueueInfo, host string, vhost string, time time.Time) error {
	var metrics []graphite.Metric = []graphite.Metric{}

	var prefix = "queues."
	var unixtime = time.Unix()

	for _,queue := range queues {
		var queuePrefix = prefix + queue.Name

		var metric = graphite.NewMetric(queuePrefix + ".messages.total", strconv.Itoa(queue.Messages), unixtime)
		metrics = append(metrics, metric)

		metric = graphite.NewMetric(queuePrefix + ".messages.ready", strconv.Itoa(queue.MessagesReady), unixtime)
		metrics = append(metrics, metric)
	}

	logrus.Info(fmt.Sprintf("Sending %d metrics to Graphite", len(metrics)))

	err := g.client.SendMetrics(metrics)

	if err != nil {
		return err
	}

	return nil
}
