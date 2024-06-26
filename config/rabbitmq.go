package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var CH *amqp.Channel

func init() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Printf("failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("failed to open a channel: %s", err)
	}

	CH = ch
	_, err = CH.QueueDeclare("person", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("fail create a queue: %s", err)
	}

}
