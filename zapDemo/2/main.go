package main

import "go.uber.org/zap"

func main() {

	url := "http://www.baidu.com"

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Debugf("Failed to fetch URL: %s", url)
	//sugar.Warnf("Failed to fetch URL: %s", url)
	//sugar.Errorf("Failed to fetch URL: %s", url)
}
