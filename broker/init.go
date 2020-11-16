package broker

import "github.com/streadway/amqp"

func RabbitMqConnection() (*amqp.Connection, error) {
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
	return connection, err
}
