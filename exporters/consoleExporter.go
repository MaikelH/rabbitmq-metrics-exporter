package exporters

import (
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/maikelh/rabbitmq-metrics-exporter/structs"
	"time"
	"strings"
)

type ConsoleExporter struct {
	Host string
	Port int

}

func NewConsoleExporter() (*ConsoleExporter, error) {
	g := new(ConsoleExporter)

	return g, nil
}

func (g *ConsoleExporter) UpdateQueues(queues []structs.Queue, host string, vhost string, time time.Time) error {
	var prefix = "queues."

	for _,queue := range queues {
		// Replace all dots in the string name by hypens, since dots dictate different metrics in graphite
		var queueName = strings.Replace(queue.Name, ".", "-", -1)

		var queuePrefix = prefix + queueName

		fmt.Printf("%s - %d\n", queuePrefix + ".messages.total", int64(queue.MessagesTotal))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.ready", int64(queue.MessagesReady))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.unacknowledged", int64(queue.MessagesUnacknowledged))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.bytes", int64(queue.MessageBytes))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.bytes-ready", int64(queue.MessageBytesReady))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.ram", int64(queue.MessagesRAM))
		fmt.Printf("%s - %d\n", queuePrefix + ".messages.persistent", int64(queue.MessagesPersistent))

		fmt.Printf("%s - %d\n", queuePrefix + ".rates.deliver", int64(queue.RateDelivered))
		fmt.Printf("%s - %d\n", queuePrefix + ".rates.deliver-get", int64(queue.RateDeliveredGet))
		fmt.Printf("%s - %d\n", queuePrefix + ".rates.deliver-noack", int64(queue.RateDeliveredNoAck))
		fmt.Printf("%s - %d\n", queuePrefix + ".rates.publish", int64(queue.RatePublished))
		fmt.Printf("%s - %d\n", queuePrefix + ".rates.redeliver", int64(queue.RateRedelivered))
	}

	logrus.Info("Sending metrics to Console")

	return nil
}