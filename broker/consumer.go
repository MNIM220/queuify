package broker

import "github.com/streadway/amqp"

type Consumer struct {
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func (c *Consumer) DefaultConsumer() (<-chan amqp.Delivery, error) {
	return c.ch.Consume(c.q.Name, "", true, false, false, false, nil)
}
