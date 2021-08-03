package pkg

import (
	"context"
	"fmt"

	"github.com/AlexDespod/testtask/shared"
	"github.com/streadway/amqp"
)

//Publisher is a type for sender
type Publisher struct {
	Conn  *amqp.Connection
	Chan  *amqp.Channel
	Queue *amqp.Queue
}

//CloseAll cancel a delivery channel and close a connection to server
func (c *Publisher) CloseAll() {
	c.Chan.Close()
	c.Conn.Close()
}

//GetPublisher return Publisher object wich contain connection , channel and queue
func GetPublisher(name string) (*Publisher, error) {

	conn, err := amqp.Dial(shared.ConnPath)

	if err != nil {
		return nil, fmt.Errorf("can't establish a connection . %v", err)
	}

	ch, err := conn.Channel()

	if err != nil {

		conn.Close()

		return nil, fmt.Errorf("can't get a channel . %v", err)
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
		return nil, fmt.Errorf("can't declare queue . %v", err)
	}

	return &Publisher{Conn: conn, Chan: ch, Queue: &q}, nil
}

//SendMessToMQ take context with Publisher and write data (message) to message broker
func SendMessToMQ(ctx context.Context, data string) error {

	pb, ok := ctx.Value("publisher").(*Publisher)

	if !ok {
		return fmt.Errorf("can't assert a type from context to Publisher")
	}

	err := pb.Chan.Publish(
		"",            // exchange
		pb.Queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	if err != nil {
		return fmt.Errorf("can't send message to broker . %v", err)
	}

	return nil
}
