package main

import (
	"net/http"
	"errors"
	"encoding/json"
	"fmt"
)

type Queues []struct {
	MessagesDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_details"`
	Messages int `json:"messages"`
	MessagesUnacknowledgedDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_unacknowledged_details"`
	MessagesUnacknowledged int `json:"messages_unacknowledged"`
	MessagesReadyDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_ready_details"`
	MessagesReady int `json:"messages_ready"`
	ReductionsDetails struct {
		Rate float64 `json:"rate"`
	} `json:"reductions_details"`
	Reductions int `json:"reductions"`
	Node string `json:"node"`
	Arguments struct {
	} `json:"arguments"`
	Exclusive bool `json:"exclusive"`
	AutoDelete bool `json:"auto_delete"`
	Durable bool `json:"durable"`
	Vhost string `json:"vhost"`
	Name string `json:"name"`
	MessageBytesPagedOut int `json:"message_bytes_paged_out"`
	MessagesPagedOut int `json:"messages_paged_out"`
	BackingQueueStatus struct {
		Mode string `json:"mode"`
		Q1 int `json:"q1"`
		Q2 int `json:"q2"`
		Delta []interface{} `json:"delta"`
		Q3 int `json:"q3"`
		Q4 int `json:"q4"`
		Len int `json:"len"`
		TargetRAMCount string `json:"target_ram_count"`
		NextSeqID int `json:"next_seq_id"`
		AvgIngressRate float64 `json:"avg_ingress_rate"`
		AvgEgressRate float64 `json:"avg_egress_rate"`
		AvgAckIngressRate float64 `json:"avg_ack_ingress_rate"`
		AvgAckEgressRate float64 `json:"avg_ack_egress_rate"`
	} `json:"backing_queue_status"`
	HeadMessageTimestamp interface{} `json:"head_message_timestamp"`
	MessageBytesPersistent int `json:"message_bytes_persistent"`
	MessageBytesRAM int `json:"message_bytes_ram"`
	MessageBytesUnacknowledged int `json:"message_bytes_unacknowledged"`
	MessageBytesReady int `json:"message_bytes_ready"`
	MessageBytes int `json:"message_bytes"`
	MessagesPersistent int `json:"messages_persistent"`
	MessagesUnacknowledgedRAM int `json:"messages_unacknowledged_ram"`
	MessagesReadyRAM int `json:"messages_ready_ram"`
	MessagesRAM int `json:"messages_ram"`
	GarbageCollection struct {
		MinorGcs int `json:"minor_gcs"`
		FullsweepAfter int `json:"fullsweep_after"`
		MinHeapSize int `json:"min_heap_size"`
		MinBinVheapSize int `json:"min_bin_vheap_size"`
		MaxHeapSize int `json:"max_heap_size"`
	} `json:"garbage_collection"`
	State string `json:"state"`
	RecoverableSlaves interface{} `json:"recoverable_slaves"`
	Consumers int `json:"consumers"`
	ExclusiveConsumerTag interface{} `json:"exclusive_consumer_tag"`
	Policy interface{} `json:"policy"`
	ConsumerUtilisation interface{} `json:"consumer_utilisation"`
	IdleSince string `json:"idle_since"`
	Memory int `json:"memory"`
}

type RabbitMQ interface {
	GetQueues(vhost string) (queues Queues, err error)
}

type RabbitMQApi struct {
	Username string
	Password string
	Host string
	Vhost string
}

func (r *RabbitMQApi) GetQueues(vhost string) (queues Queues, err error) {
	url := fmt.Sprintf("http://%s:15672/api/queues", r.Host)

	response, err := http.Get(url)

	if response.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("Got statuscode: %d", response.StatusCode))
	}

	if err != nil {
		return nil, errors.New("Something went wrong during http request.")
	}

	var q Queues

	err = json.NewDecoder(response.Body).Decode(&q)

	return q,nil
}