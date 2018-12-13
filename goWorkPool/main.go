package main

import (
	"time"

	"./workerpool"
)

func main() {
	d := workerpool.NewDispatcher(20000)
	d.Start()

	p := workerpool.NewProducer(1000000)
	p.Run()

	time.Sleep(100 * time.Second)
}
