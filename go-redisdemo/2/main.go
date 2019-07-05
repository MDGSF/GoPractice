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

func ExampleClient() {
	log.Info("ExampleClient start")

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		log.Error("err = %v", err)
		return
	}

	val, err := client.Get("key").Result()
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		log.Error("err = %v", err)
		return
	} else {
		fmt.Println("key2", val2)
	}

	fmt.Println()
}

func main() {
	fmt.Println("vim-go")
	ExampleNewClient()
	for {
		ExampleClient()
		time.Sleep(time.Second)
	}
}
