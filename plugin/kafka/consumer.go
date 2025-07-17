package kafka

import (
	sctx "blog-tech/plugin"
	"context"

	"github.com/segmentio/kafka-go"
)

type MessageHandler func(key []byte, value []byte) error

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(config *Config, topic string) *Consumer {
	return &Consumer{
		reader: config.NewReader(topic),
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func (c *Consumer) Consume(ctx context.Context, handler MessageHandler) {
	logger := sctx.GlobalLogger().GetLogger("service")

	for {
		msg, err := c.reader.ReadMessage(ctx)

		if err != nil {
			logger.Error("Error reading message from Kafka: %v", err)
			continue
		}

		if err := handler(msg.Key, msg.Value); err != nil {
			logger.Error("Error processing message: %v", err)
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			logger.Error("Error committing message: %v", err)
		}
	}
}
