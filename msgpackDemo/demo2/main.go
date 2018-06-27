package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type TMsg struct {
	Student1 TStudent  `msgpack:"student1"`
	Student2 *TStudent `msgpack:"student2"`
}

type TMsg2 struct {
	Student1 TStudent `msgpack:"student1"`
	Student2 TStudent `msgpack:"student2"`
}

type TStudent struct {
	ID   string `msgpack:"id"`
	Name string `msgpack:"name"`
	Age  int    `msgpack:"age"`
}

func main() {

	msg := &TMsg{
		Student1: TStudent{
			ID:   "ID1",
			Name: "Name1",
			Age:  1,
		},
		Student2: &TStudent{
			ID:   "ID2",
			Name: "Name2",
			Age:  2,
		},
	}

	b, err := msgpack.Marshal(msg)

	s := &TMsg2{}
	err = msgpack.Unmarshal(b, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(s.Student1)
	fmt.Println(s.Student2)
}
