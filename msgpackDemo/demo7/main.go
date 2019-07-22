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
	test3()
}

func test3() {
	data := []byte{131, 163, 97, 103, 101, 203, 64, 195, 136, 0, 0, 0, 0, 0, 163, 115, 117, 98, 131, 161, 97, 203, 64, 89, 0, 0, 0, 0, 0, 0, 161, 99, 129, 168, 108, 97, 110, 103, 117, 97, 103, 101, 164, 82, 117, 115, 116, 161, 98, 194, 164, 110, 97, 109, 101, 169, 104, 117, 97, 110, 103, 106, 105, 97, 110}

	var output interface{}
	err := msgpack.Unmarshal(data, &output)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	fmt.Println("output = ", output)
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
