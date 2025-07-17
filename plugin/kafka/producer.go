package kafka

import (
	sctx "blog-tech/plugin"
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(config *Config, topic string) *Producer {
	return &Producer{
		writer: config.NewWriter(topic),
	}
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

func (p *Producer) Produce(ctx context.Context, key string, value interface{}) error {
	logger := sctx.GlobalLogger().GetLogger("service")

	data, err := json.Marshal(value)

	if err != nil {
		logger.Error("Failed to marshal message: %v", err)
		return err
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	maxRetries := 3

	for i := range maxRetries {
		produceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		err = p.writer.WriteMessages(produceCtx, msg)
		cancel()
		if err == nil {
			logger.Info("Message sent to Kafka successfully - key: %s", key)
			return nil
		}

		logger.Warn("Failed to write message to Kafka (attempt %d/%d): %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	logger.Error("Failed to write message to Kafka after %d attempts", maxRetries)

	return nil
}
