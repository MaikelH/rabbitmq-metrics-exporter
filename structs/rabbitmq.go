package structs

type Queue struct {
	Name string
	Vhost string
	Node string
	Durable bool
	AutoDelete bool
	MessagesTotal 		int64
	MessagesReady 		int64
	MessagesUnacknowledged 	int64
	MessageBytesReady	int64
	MessagesRAM		int64
	MessagesPersistent	int64

	MessageBytes 		int64

	RateDelivered      	int64
	RateDeliveredGet   	int64
	RateDeliveredNoAck 	int64
	RatePublished      	int64

	RateRedelivered 	int64
}
