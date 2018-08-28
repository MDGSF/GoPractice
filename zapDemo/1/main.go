package main

import (
	"io/ioutil"
	"time"

	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("main start")
	logger.Info("failed to fetch URL",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	_, err := ioutil.ReadFile("test")
	if err != nil {
		logger.Error("failed to read file", zap.String("filename", "test"), zap.Any("err", err))
	}

	success := false
	if !success {
		logger.Error("failed to connect to server", zap.String("err", "invalid network"))
	}

	logger.Info("main end")
}
