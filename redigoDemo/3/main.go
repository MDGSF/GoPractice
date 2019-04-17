package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var pool *redis.Pool

func init() {
	pool = newPool("127.0.0.1:6379")
}

func main() {
	for i := 0; i < 1; i++ {
		go func() {
			conn := pool.Get()
			defer conn.Close()

			psc := redis.PubSubConn{Conn: conn}
			psc.Subscribe("example")
			for {
				switch v := psc.Receive().(type) {
				case redis.Message:
					fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
				case redis.Subscription:
					fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
				case error:
					return
				}
			}
		}()
	}

	conn := pool.Get()
	defer conn.Close()
	for {
		var s string
		fmt.Scanln(&s)
		_, err := conn.Do("PUBLISH", "example", s)
		if err != nil {
			fmt.Println("publish err:", err)
			return
		}
	}
}
