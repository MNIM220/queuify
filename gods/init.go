package gods

import (
	"queuify/broker"
)

func WorkerGods() {
	publisherConnection, err := broker.RabbitMqConnection()
	if err != nil {
		panic(err)
	}
	concumerConnection, err := broker.RabbitMqConnection()
	if err != nil {
		panic(err)
	}

}
