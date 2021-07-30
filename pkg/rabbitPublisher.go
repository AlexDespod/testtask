package pkg

import (
	"context"
	"log"

	"github.com/AlexDespod/testtask/shared"
	"github.com/streadway/amqp"
)

func GetPublisher() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {

	conn, err := amqp.Dial(shared.ConnPath)

	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func SendMessToMQ(ctx context.Context, filePath string) {

	pb, ok := ctx.Value("publisher").(shared.Publisher)
	if !ok {
		return
	}
	err := pb.Chan.Publish(
		"",            // exchange
		pb.Queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(filePath),
		})
	failOnError(err, "cant send")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
