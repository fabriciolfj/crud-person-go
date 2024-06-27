package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Fatalf("Falha ao conectar ao RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir um canal: %v", err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"person", // queue
		"",       // consumer
		false,    // auto-ack (agora é false para acks manuais)
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	if err != nil {
		log.Fatalf("Falha ao registrar o consumidor: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Mensagem recebida: %s", d.Body)

			var person Person
			err := json.Unmarshal(d.Body, &person)
			if err != nil {
				log.Printf("Erro ao decodificar JSON: %v", err)
				d.Nack(false, true) // Rejeita a mensagem e recoloca na fila
				continue
			}

			// Processa a mensagem
			log.Printf("Pessoa processada: %+v", person)

			// Acknowledge a mensagem
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Aguardando mensagens. Para sair pressione CTRL+C")
	<-forever
}
