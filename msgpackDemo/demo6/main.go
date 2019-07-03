package main

import (
	"fmt"

	"github.com/MDGSF/utils/log"
	"github.com/vmihailenco/msgpack"
)

func main() {
	test2()
}

func showBytes(data []byte) {
	for _, v := range data {
		fmt.Printf("%02X ", v)
	}
	fmt.Println()
}

func test2() {
	data, err := msgpack.Marshal(nil, 1, 2, nil, 200)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	showBytes(data)
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
