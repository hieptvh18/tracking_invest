package main

import (
	"context"
	"fmt"
	"time"

	"api_golang/internal/kafka"
)

// command test send data to Kafka topic "test-topic": using producer
func main() {

	brokers := []string{
		"localhost:19092",
		"localhost:29092",
		"localhost:39092",
	}

	producer := kafka.NewProducer(brokers, "test-topic")
	defer producer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := fmt.Sprintf("Hello at %s", time.Now())

	err := producer.Send(ctx, "new key", msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Message sent:", msg)
}