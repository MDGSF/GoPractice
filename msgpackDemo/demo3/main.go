package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type TMsg struct {
	Name  string   `msgpack:"name"`
	Image []string `msgpack:"image,omitempty"`
}

type TMsg2 struct {
	Name  string   `msgpack:"name"`
	Image []string `msgpack:"image"`
}

func main() {
	test1()
	test2()
	test3()
	testMsg2_1()
	testMsg2_2()
	testMsg2_3()
}

func testMsg2_3() {
	msg := &TMsg2{Name: "testMsg2_2"}
	msg.Image = make([]string, 0)
	msg.Image = append(msg.Image, "I'm image")

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}

func testMsg2_2() {
	msg := &TMsg2{Name: "testMsg2_2"}
	msg.Image = nil

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}

func testMsg2_1() {
	msg := &TMsg2{Name: "testMsg2_1"}
	msg.Image = make([]string, 0)

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}

func test3() {
	msg := &TMsg{Name: "test3"}
	msg.Image = make([]string, 0)
	msg.Image = append(msg.Image, "I'm image")

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}

func test2() {
	msg := &TMsg{Name: "test2"}
	msg.Image = nil

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}

func test1() {
	msg := &TMsg{Name: "test1"}
	msg.Image = make([]string, 0)

	b, _ := msgpack.Marshal(msg)
	fmt.Println(string(b))
}
