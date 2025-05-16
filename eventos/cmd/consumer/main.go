package main

import (
	"eventos/pkg/rabbitmq"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cs, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer cs.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(cs, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
