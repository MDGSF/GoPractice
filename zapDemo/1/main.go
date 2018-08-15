package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	logger.Info("failed to fetch URL",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
