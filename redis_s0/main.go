package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MDGSF/utils/log"
	"github.com/gomodule/redigo/redis"
)

var (
	C1AuthKey = "c1_auth_keys"
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

func GetAllC1Key() {
	conn := pool.Get()
	defer conn.Close()

	result, err := redis.Strings(conn.Do("hgetall", C1AuthKey))
	if err != nil {
		log.Error("redis get c1id failed, err = %v", err)
		return
	}

	fmt.Printf("var data = []string{\n")
	for i := 0; i < len(result); i += 2 {
		fmt.Printf("\"%v\", \"%v\",\n", result[i], result[i+1])
	}
	fmt.Printf("}")
}

var data = []string{}

func restore() {
	for i := 0; i < len(data); i += 2 {
		setC1Key(data[i], data[i+1])
	}
}

func setC1Key(c1id, key string) {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("hset", C1AuthKey, c1id, key)
	if err != nil {
		log.Error("redis get c1id failed, err = %v", err)
		return
	}
}

func delC1Key(c1id string) {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("hdel", C1AuthKey, c1id)
	if err != nil {
		log.Error("redis get c1id failed, err = %v", err)
		return
	}
}

func deleteC1Name() {
	var name string
	flag.StringVar(&name, "name", "", "c1 name to delete from redis")

	flag.Parse()

	if len(name) == 0 {
		log.Error("empty name")
		os.Exit(0)
	}

	delC1Key(name)
}

func main() {
	pool = newPool("127.0.0.1:6379")

	//GetAllC1Key()
	// restore()

	deleteC1Name()
}
