package main

import (
	"context"
	"os"

	"github.com/dlbarduzzi/demo/internal/logging"
)

func main() {
	logger := logging.NewLogger("dev", "debug")
	logger.Info("Hello from dev!")

	logger = logging.NewLogger("prod", "info")
	logger.Info("Hello from prod!")

	logger = logging.NewLoggerFromEnv()
	logger.Info("Hello from env mode!")

	logger = logging.NewLoggerFromEnv().With("app", "demo")

	ctx := context.Background()
	ctx = logging.LoggerWithContext(ctx, logger)

	if err := start(ctx); err != nil {
		logger.Error(err.Error())
		os.Exit(2)
	}
}

func start(ctx context.Context) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Info("calling from start")
	return nil
}
