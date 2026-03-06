package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"api_golang/internal/kafka"
)

func main() {
	brokers := []string{
		"localhost:19092",
		"localhost:29092",
		"localhost:39092",
	}

	consumer :=kafka.NewConsumer(brokers, "test-topic")
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		cancel()
	}()

	fmt.Println("Consumer listening on topic test-topic (Ctrl+C to stop)...")

	err := consumer.Read(ctx, func(key, value string) error {
		fmt.Printf("key=%s value=%s\n", key, value)
		return nil
	})
	if err != nil && ctx.Err() == nil {
		panic(err)
	}
}
