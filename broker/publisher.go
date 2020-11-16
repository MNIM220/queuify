package broker

import "github.com/streadway/amqp"

type Publisher struct {
	ch amqp.Channel
	q  amqp.Queue
}

func (p *Publisher) DefaultTextPublisher(data []byte) error {
	err := p.ch.Publish("", p.q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	if err != nil {
		return err
	}
	return nil
}
