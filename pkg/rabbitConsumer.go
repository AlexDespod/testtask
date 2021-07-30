package pkg

import (
	"github.com/AlexDespod/testtask/shared"
	"github.com/streadway/amqp"
)

func GetConsumer(name string) (<-chan amqp.Delivery, *amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(shared.ConnPath)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	return msgs, conn, ch
}
