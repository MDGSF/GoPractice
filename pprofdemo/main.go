package main

import (
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		time.Sleep(time.Second)
	}
}
