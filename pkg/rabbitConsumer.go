package pkg

import (
	"github.com/AlexDespod/testtask/shared"
	"github.com/streadway/amqp"
)

type Consumer struct {
	Delivery <-chan amqp.Delivery
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Name     string
}

//CloseAll cancel a delivery channel and close a connection to server
func (c *Consumer) CloseAll() {

	c.Channel.Close()
	c.Conn.Close()
}

// GetConsumer return consumer for "name from args" queue . You should to close a amqp.Connection and amqp.Channel
func GetConsumer(name string) (*Consumer, error) {
	conn, err := amqp.Dial(shared.ConnPath)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		conn.Close()
		return nil, err
	}
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		conn.Close()
		ch.Close()
		//ch.Cancel(name, false)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		conn.Close()
		ch.Close()
		// ch.Cancel(name, false)
		return nil, err
	}

	return &Consumer{msgs, conn, ch, name}, nil
}
