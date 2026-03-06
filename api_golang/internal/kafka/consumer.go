package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const consumerGroupID = "test-consumer-group"

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: consumerGroupID,
		}),
	}
}

func (c *Consumer) Read(ctx context.Context, handler func(key, value string) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled {
				return nil
			}
			return fmt.Errorf("read message: %w", err)
		}
		if handler != nil {
			err = handler(string(msg.Key), string(msg.Value))
			if err != nil {
				log.Printf("handler error: %v", err)
			}
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}