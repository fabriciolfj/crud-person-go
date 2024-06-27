package service

import (
	"encoding/json"
	rabbitmq "github.com/person/config"
	"github.com/person/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func SendMessage(p *model.Person) {
	value, err := json.Marshal(p)

	if err != nil {
		log.Fatalf("fail encode to json: %v", err)
		return
	}

	err = rabbitmq.CH.Publish("person", "person", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        value,
	})

	if err != nil {
		log.Fatalf("failed to publish message: %s", err)
		return
	}

	log.Printf("message send successfully, %v\n", string(value))
}
