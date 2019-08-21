package main

import (
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {
	t1 := time.NewTimer(5 * time.Second)

	t2 := time.NewTimer(10 * time.Second)

	t3 := time.NewTicker(60 * time.Second)

	for {
		select {
		case <-t1.C:
			log.Info("t1")
			if !t2.Stop() {
				<-t2.C
			}
			ret := t2.Reset(time.Second * 2)
			log.Info("ret = %v", ret)
		case <-t2.C:
			log.Info("t2")
			break
		case <-t3.C:
			log.Info("t3")
		}
	}
}
