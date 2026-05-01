//go:build external

package main

import "go.uber.org/zap"

func initLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
	return logger
}

func main() {
	logger := initLogger()
	defer logger.Sync()
	logger.Info("Logger initialized")
}
