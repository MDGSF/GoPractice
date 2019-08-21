package main

import (
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {
	t1 := time.NewTimer(10 * time.Second)

	t2 := time.NewTimer(5 * time.Second)

	t3 := time.NewTimer(20 * time.Second)

	for {
		select {
		case <-t1.C:
			log.Info("t1")
		case <-t2.C:
			log.Info("t2")
			t1.Stop()
		case <-t3.C:
			log.Info("t3")
		}
	}
}
