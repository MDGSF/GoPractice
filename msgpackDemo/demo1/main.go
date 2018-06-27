package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

func main() {
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
