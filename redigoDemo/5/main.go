package main

import (
	"fmt"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/gomodule/redigo/redis"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:         3,
		MaxActive:       0,
		IdleTimeout:     0,
		MaxConnLifetime: 0,
		Dial:            func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var pool *redis.Pool

func init() {
	pool = newPool("127.0.0.1:6379")
}

func main() {
	for {
		log.Println("for start")
		conn := pool.Get()
		if conn.Err() != nil {
			log.Error("pool.Get() failed, err = %v", conn.Err())
			time.Sleep(time.Second)
			continue
		}

		psc := redis.PubSubConn{Conn: conn}
		err := psc.Subscribe("example")
		if err != nil {
			log.Error("psc subscribe failed, err = %v", conn.Err())
			time.Sleep(time.Second)
			continue
		}

		for conn.Err() == nil {
			switch v := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Printf("find an error: %v\n", v)
			}
		}

		if conn.Err() != nil {
			log.Printf("conn.Err() = %v\n", conn.Err())
		}

		conn.Close()
	}
}
