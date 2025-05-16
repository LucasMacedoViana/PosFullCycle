package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	//Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Receivid from RabbitMQ: ID: %d - %s\n", msg.ID, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Receivid from Kafka: ID: %d - %s\n", msg.ID, msg.Msg)
		case <-time.After(3 * time.Second):
			println("Timeout")
		}
	}
}
