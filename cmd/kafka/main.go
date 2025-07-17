package main

import (
	"blog-tech/cmd"
	"blog-tech/common"
	sctx "blog-tech/plugin"
	"blog-tech/plugin/kafka/consumer"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger := sctx.GlobalLogger().GetLogger("consumer")
	logger.Info("Starting consumer application...")

	time.Sleep(time.Second * 2)

	// Dùng shared service context
	serviceCtx := cmd.NewServiceCtx()

	if err := serviceCtx.Load(); err != nil {
		logger.Fatalf("Failed to load service context: %v", err)
	}

	// Create consumer engine
	engine := consumer.NewConsumerEngine(serviceCtx)

	// Register event handlers
	engine.Register(common.ArticleLikeTopic, consumer.HandleArticleLikeEvent)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start consumer engine
	if err := engine.Start(ctx); err != nil {
		logger.Fatalf("Failed to start consumer engine: %v", err)
	}

	logger.Infof("Consumer engine started with topics: %v", engine.GetActiveTopics())

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Graceful shutdown
	logger.Info("Shutting down consumer engine...")
	cancel()

	if err := engine.Stop(); err != nil {
		logger.Errorf("Error stopping consumer engine: %v", err)
	}

	logger.Info("Consumer engine shutdown complete")
}
