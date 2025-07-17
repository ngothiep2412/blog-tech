package consumer

import (
	sctx "blog-tech/plugin"
	"blog-tech/plugin/kafka"
	"context"
	"sync"
)

type MessageHandler func(serviceCtx sctx.ServiceContext, key []byte, value []byte) error

type ConsumerConfig struct {
	Topic   string
	Handler MessageHandler
}

type ConsumerEngine struct {
	serviceCtx  sctx.ServiceContext
	kafkaConfig *kafka.Config
	consumers   map[string]*kafka.Consumer
	configs     []ConsumerConfig
	mu          sync.RWMutex
}

func NewConsumerEngine(serviceCtx sctx.ServiceContext) *ConsumerEngine {
	return &ConsumerEngine{
		serviceCtx:  serviceCtx,
		kafkaConfig: kafka.NewConfig(),
		consumers:   make(map[string]*kafka.Consumer),
		configs:     make([]ConsumerConfig, 0),
	}
}

func (engine *ConsumerEngine) Register(topic string, handler MessageHandler) *ConsumerEngine {
	engine.mu.Lock()
	defer engine.mu.Unlock()

	logger := sctx.GlobalLogger().GetLogger("service")

	engine.configs = append(engine.configs, ConsumerConfig{
		Topic:   topic,
		Handler: handler,
	})

	logger.Info("Registered consumer for topic: %s", topic)
	return engine
}

func (engine *ConsumerEngine) Start(ctx context.Context) error {
	engine.mu.Lock()
	defer engine.mu.Unlock()

	logger := sctx.GlobalLogger().GetLogger("service")

	for _, config := range engine.configs {
		consumer := kafka.NewConsumer(engine.kafkaConfig, config.Topic)
		engine.consumers[config.Topic] = consumer

		go func(topic string, handler MessageHandler, consumer *kafka.Consumer) {
			logger.Info("Starting consumer for topic: %s", topic)

			consumer.Consume(ctx, func(key []byte, value []byte) error {
				return handler(engine.serviceCtx, key, value)
			})
		}(config.Topic, config.Handler, consumer)
	}

	logger.Info("Started %d consumers", len(engine.configs))
	return nil
}

func (engine *ConsumerEngine) Stop() error {
	engine.mu.Lock()
	defer engine.mu.Unlock()

	logger := sctx.GlobalLogger().GetLogger("service")

	for topic, consumer := range engine.consumers {
		if err := consumer.Close(); err != nil {
			logger.Errorf("Error closing consumer for topic %s: %v", topic, err)
		} else {
			logger.Info("Closed consumer for topic: %s", topic)
		}
	}

	engine.consumers = make(map[string]*kafka.Consumer)
	return nil
}

func (engine *ConsumerEngine) GetActiveTopics() []string {
	engine.mu.RLock()
	defer engine.mu.RUnlock()

	topics := make([]string, 0, len(engine.consumers))
	for topic := range engine.consumers {
		topics = append(topics, topic)
	}

	return topics
}
