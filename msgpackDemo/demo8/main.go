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

	var temp interface{}
	err = msgpack.Unmarshal(b, &temp)
	if err != nil {
		panic(err)
	}
	fmt.Println(temp)

	tempdata, err := msgpack.Marshal(temp)
	if err != nil {
		panic(err)
	}

	s := &Student{}
	err = msgpack.Unmarshal(tempdata, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
