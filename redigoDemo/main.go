/*
hset(key, field, value)：向名称为key的hash中添加元素field
hget(key, field)：返回名称为key的hash中field对应的value
hmget(key, (fields))：返回名称为key的hash中field i对应的value
hmset(key, (fields))：向名称为key的hash中添加元素field
hincrby(key, field, integer)：将名称为key的hash中field的value增加integer
hexists(key, field)：名称为key的hash中是否存在键为field的域
hdel(key, field)：删除名称为key的hash中键为field的域
hlen(key)：返回名称为key的hash中元素个数
hkeys(key)：返回名称为key的hash中所有键
hvals(key)：返回名称为key的hash中所有键对应的value
hgetall(key)：返回名称为key的hash中所有的键（field）及其对应的value

*/
package main

import (
	"fmt"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/gomodule/redigo/redis"
)

var c redis.Conn

func main() {
	start()
	test8()
}

func start() {
	var err error
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis error", err)
		return
	}
}

// set get
func test1() {
	_, err := c.Do("SET", "mykey", "huangjian")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey: %v", username)
}

// test set expire time.
func test2() {
	_, err := c.Do("SET", "mykey", "huangjian", "EX", 5)
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey 1: %v", username)

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey 2: %v", username)
}

// exists
func test3() {
	_, err := c.Do("SET", "mykey", "huangjian")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	is_key_exist, err := redis.Int64(c.Do("EXISTS", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("is_key_exist = %v", is_key_exist)
}

// delete
func test4() {
	_, err := c.Do("SET", "mykey", "huangjian")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey: %v", username)

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		log.Error("redis delete failed, err = %v", err)
		return
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Error("redis get failed, err = %v", err)
		return
	}

	log.Info("get mykey: %v", username)
}

// hset
func test5() {
	_, err := c.Do("HSET", "student", "id", "1111111111", "name", "huangjian", "age", 12)
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	id, err := redis.String(c.Do("HGET", "student", "id"))
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}
	name, err := redis.String(c.Do("HGET", "student", "name"))
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}
	age, err := redis.Int(c.Do("HGET", "student", "age"))
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}
	log.Info("id = %v", id)
	log.Info("name = %v", name)
	log.Info("age = %v", age)
}

// hgetall
func test6() {
	_, err := c.Do("HSET", "student", "id", "1111111111", "name", "huangjian", "age", 12)
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	result, err := redis.Values(c.Do("hgetall", "student"))
	if err != nil {
		log.Error("redis hgetall failed, err = %v", err)
		return
	}
	for _, v := range result {
		log.Info("%v", string(v.([]byte)))
	}
}

// hgetall
func test7() {
	_, err := c.Do("HSET", "student", "id", "1111111111", "name", "huangjian", "age", 12)
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	result, err := redis.Strings(c.Do("hgetall", "student"))
	if err != nil {
		log.Error("redis hgetall failed, err = %v", err)
		return
	}
	for _, v := range result {
		log.Info("%v", v)
	}
}

// send
func test8() {
	err := c.Send("sadd", "studentSet", "huangjian")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}
	err = c.Send("sadd", "studentSet", "huangping")
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}
	_, err = c.Do("HSET", "student::1111", "id", "1111111111", "name", "huangjian", "age", 12)
	if err != nil {
		log.Error("redis set failed, err = %v", err)
		return
	}

	c.Flush()
}
