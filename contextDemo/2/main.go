package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func doStuff() {
	ctx, cancel := context.WithCancel(context.Background())

	i := 0
	for {
		time.Sleep(time.Second)

		i++
		if i > 3 {
			cancel()
		}

		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)

	go doStuff()
	time.Sleep(10 * time.Second)

	logg.Printf("down")
}
