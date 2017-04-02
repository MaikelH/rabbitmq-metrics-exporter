package structs

type Queue struct {
	Name string
	Vhost string
	Node string
	Durable bool
	AutoDelete bool
	MessagesTotal int64
	MessagesReady int64
	MessagesUnacknowledged int64

	MessageBytes int64

	RateDelivered int64
	RatePublished int64
}
