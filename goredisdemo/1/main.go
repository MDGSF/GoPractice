package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		MinIdleConns:       0,
		MaxConnAge:         10 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 1 * time.Second,
	})
	defer client.Close()

	pubsub := client.Subscribe("example")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
