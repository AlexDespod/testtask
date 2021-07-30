package shared

import "github.com/streadway/amqp"

const ConnPath = "amqp://guest:guest@localhost:5672/"

type Publisher struct {
	Chan  *amqp.Channel
	Queue *amqp.Queue
}
