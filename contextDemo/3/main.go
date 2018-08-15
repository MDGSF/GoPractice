package main

import (
	"context"
	"time"

	"github.com/MDGSF/utils/log"
)

func someHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doStuff(ctx)

	time.Sleep(10 * time.Second)
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			log.Info("deadline set")
			if time.Now().After(deadline) {
				log.Info("%v", ctx.Err().Error())
			}
		}

		select {
		case <-ctx.Done():
			log.Info("done")
			return
		default:
			log.Info("work")
		}
	}
}

func main() {
	log.SetLevel(log.VerboseLevel)
	someHandler()
	log.Info("down")
}
