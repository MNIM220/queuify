package broker

import "github.com/streadway/amqp"

type Publisher struct {
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func (p *Publisher) DefaultTextPublisher(data []byte) error {
	err := p.Channel.Publish("", p.Queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	if err != nil {
		return err
	}
	return nil
}
