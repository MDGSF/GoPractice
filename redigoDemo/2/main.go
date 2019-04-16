package main

import (
	"time"

	"github.com/MDGSF/utils/log"
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
	for {
		test1()
		time.Sleep(time.Second * 3)
	}
}

func test1() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "mykey", "huangjian")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey: %v", username)
}
