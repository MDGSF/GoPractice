package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type TMsg struct {
	Name  string   `msgpack:"name"`
	Image []string `msgpack:"image,omitempty"`
}

func main() {
	fmt.Println("vim-go")

	msg1 := TMsg{
		Name: "msg1",
	}

	msg2 := TMsg{
		Name: "msg2",
	}

	m := make(map[int]TMsg)
	m[1] = msg1
	m[2] = msg2

	b, err := msgpack.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(b))

	m2 := make(map[int]TMsg)
	err = msgpack.Unmarshal(b, &m2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m2)
}
