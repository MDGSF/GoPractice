package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/vmihailenco/msgpack"
)

func main() {
	test2()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func showBytes(data []byte) {
	for _, v := range data {
		fmt.Printf("0x%02X, ", v)
	}
	fmt.Println()
}

func test2() {
	input := false
	fmt.Println("input = ", input)
	data, err := msgpack.Marshal(input)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	showBytes(data)

	var output interface{}
	err = msgpack.Unmarshal(data, &output)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	fmt.Println("output = ", output)
}

func test1() {
	type Student struct {
		ID   string `msgpack:"id"`
		Name string `msgpack:"name"`
		Age  int    `msgpack:"age"`
	}

	b, err := msgpack.Marshal(&Student{
		ID:   "ID123",
		Name: "Name123",
		Age:  11,
	})

	/*
		msgpack can use like json, just get data you want.
	*/
	type StudentRecv struct {
		NewAge int    `msgpack:"age"`
		NewID  string `msgpack:"id"`
	}
	s := &StudentRecv{}
	err = msgpack.Unmarshal(b, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
