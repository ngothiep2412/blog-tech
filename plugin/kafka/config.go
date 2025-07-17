package kafka

import (
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	Brokers []string
	GroupID string
}

func NewConfig() *Config {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	if len(brokers) == 0 || (len(brokers) == 1 && brokers[0] == "") {
		brokers = []string{"localhost:9092"}
	}

	return &Config{
		Brokers: brokers,
		GroupID: os.Getenv("KAFKA_GROUP_ID"),
	}
}

func (c *Config) NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(c.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
}

func (c *Config) NewReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        c.Brokers,
		GroupID:        c.GroupID,
		Topic:          topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
	})
}
