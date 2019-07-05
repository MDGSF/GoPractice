package main

import (
	"fmt"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/go-redis/redis"
)

var client *redis.Client

func ExampleNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func ExamplePubSub() {
	for {

		pubsub := client.Subscribe("example")

		for {
			msgi, err := pubsub.Receive()
			if err != nil {
				log.Error("err = %v", err)
				break
			}

			switch msg := msgi.(type) {
			case *redis.Subscription:
				log.Info("subscribed to %v", msg.Channel)
			case *redis.Message:
				log.Info("received %v from %v", msg.Payload, msg.Channel)
			default:
				log.Error("unreached")
				break
			}
		}

		pubsub.Close()
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("vim-go")
	ExampleNewClient()
	ExamplePubSub()
}
