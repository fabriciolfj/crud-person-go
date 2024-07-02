package rabbitmq

import (
	"github.com/magiconair/properties"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var CH *amqp.Channel

func init() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	host := os.Getenv("RABBITMQ_HOST")
	if host == "" {
		host = p.MustGetString("rabbitmq_url")
	}

	logrus.Info(host)
	conn, err := amqp.Dial(host)
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
