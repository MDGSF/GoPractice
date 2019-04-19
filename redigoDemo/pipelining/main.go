package main

import (
	"fmt"
	"time"

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

func withoutPipeLining() {
	start := time.Now()
	defer func() {
		fmt.Printf("withoutPipeLining elpse = %v\n", time.Now().Sub(start))
	}()

	conn := pool.Get()
	defer conn.Close()

	for i := 0; i < 10000; i++ {
		conn.Do("set", "k1", 1)
	}
}

func withPipeLining() {
	start := time.Now()
	defer func() {
		fmt.Printf("withPipeLining elpse = %v\n", time.Now().Sub(start))
	}()

	conn := pool.Get()
	defer conn.Close()

	for i := 0; i < 10000; i++ {
		conn.Send("set", "k1", 1)
	}
	conn.Do("EXEC")
}

func main() {
	withoutPipeLining()
	withPipeLining()
}
