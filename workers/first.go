package workers

import (
	"fmt"
	"queuify/broker"

	"github.com/streadway/amqp"
)

type Utils struct {
	PublisherConnection *amqp.Connection
	ConsumerConnection  *amqp.Connection
}

func (u *Utils) First(entrance []string, exit []string) {
	for _, e := range entrance {
		go func(e string) {
			consumerChannel, err := u.ConsumerConnection.Channel()
			if err != nil {
				panic(err)
			}
			consumerQueue, err := consumerChannel.QueueDeclare(e, true, false, false, false, nil)
			if err != nil {
				panic(err)
			}
			consumer := broker.Consumer{
				Queue:   &consumerQueue,
				Channel: consumerChannel,
			}

			publisherChannel, err := u.PublisherConnection.Channel()
			if err != nil {
				panic(err)
			}
			publisherQueue, err := publisherChannel.QueueDeclare(exit[0], true, false, false, false, nil)
			if err != nil {
				panic(err)
			}

			publisher := broker.Publisher{
				Queue:   &publisherQueue,
				Channel: publisherChannel,
			}
			msgs, err := consumer.DefaultConsumer()
			if err != nil {
				panic(err)
			}
			for msg := range msgs {
				fmt.Printf("hello with my sweet message %s\n", string(msg.Body))
				err := publisher.DefaultTextPublisher([]byte("hi dude whats up"))
				if err != nil {
					panic(err)
				}
			}
		}(e)
	}
}
