package exporters

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/cactus/go-statsd-client/statsd"
	"strconv"
	"fmt"
	"strings"
	"github.com/spf13/viper"
	"github.com/maikelh/rabbitmq-metrics-exporter/structs"
)

type StatsDExporter struct {
	Host string
	Port int
	client statsd.Statter
}

func NewStatsDExporter() (*StatsDExporter, error) {
	g := new(StatsDExporter)
	g.Host = viper.GetString("exporter.host")
	g.Port = viper.GetInt("exporter.port")

	err := g.setupStatsD()

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *StatsDExporter) setupStatsD() error  {
	//stats, err := statsd.NewBufferedClient(g.Host + ":" + strconv.Itoa(g.Port), "rabbitmq", 300*time.Millisecond ,0)
	logrus.Info(fmt.Sprintf("Setting StatsD host to: " + g.Host + ":" + strconv.Itoa(g.Port)))

	stats, err := statsd.NewClient(g.Host + ":" + strconv.Itoa(g.Port), "rabbitmq")

	if err != nil {
		return err
	}

	g.client = stats

	return nil
}

func (g *StatsDExporter) UpdateQueues(queues []structs.Queue, host string, vhost string, time time.Time) error {
	var prefix = "queues."

	var samplingRate = float32(1.0)

	for _,queue := range queues {
		// Replace all dots in the string name by hypens, since dots dictate different metrics in graphite
		var queueName = strings.Replace(queue.Name, ".", "-", -1)

		var queuePrefix = prefix + queueName

		g.client.Gauge(queuePrefix + ".messages.total", int64(queue.MessagesTotal), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.ready", int64(queue.MessagesReady), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.unacknowledged", int64(queue.MessagesUnacknowledged), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.bytes", int64(queue.MessageBytes), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.bytes-ready", int64(queue.MessageBytesReady), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.ram", int64(queue.MessagesRAM), samplingRate)
		g.client.Gauge(queuePrefix + ".messages.persistent", int64(queue.MessagesPersistent), samplingRate)

		g.client.Gauge(queuePrefix + ".rates.deliver", int64(queue.RateDelivered), samplingRate)
		g.client.Gauge(queuePrefix + ".rates.deliver-get", int64(queue.RateDeliveredGet), samplingRate)
		g.client.Gauge(queuePrefix + ".rates.deliver-noack", int64(queue.RateDeliveredNoAck), samplingRate)
		g.client.Gauge(queuePrefix + ".rates.publish", int64(queue.RatePublished), samplingRate)
		g.client.Gauge(queuePrefix + ".rates.redeliver", int64(queue.RateRedelivered), samplingRate)
	}

	logrus.Info("Sending metrics to StatsD")

	return nil
}
