package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// connect to RabbitMQ Server
	conn, err := amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
