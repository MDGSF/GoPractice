package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

func show(data []byte) {
	for _, v := range data {
		fmt.Printf("0x%02x ", v)
	}
	fmt.Println()

	for _, v := range data {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	fmt.Println()
}

func testStruct() {
	type Student struct {
		ID string `msgpack:"id"`
		A  int    `msgpack:"a"`
	}

	b, _ := msgpack.Marshal(&Student{
		ID: "123",
		A:  256,
	})

	show(b)
}

func testBooleanTrue() {
	i := true
	b, _ := msgpack.Marshal(&i)
	fmt.Println(true)
	show(b)
}

func testBooleanFalse() {
	i := false
	b, _ := msgpack.Marshal(&i)
	fmt.Println(false)
	show(b)
}

func testInt1() {
	i := 1
	b, _ := msgpack.Marshal(&i)
	fmt.Println(1)
	show(b)
}

func testInt10() {
	i := 10
	b, _ := msgpack.Marshal(&i)
	fmt.Println(10)
	show(b)
}

func testInt31() {
	i := 31
	b, _ := msgpack.Marshal(&i)
	fmt.Println(31)
	show(b)
}

func testInt32() {
	i := 32
	b, _ := msgpack.Marshal(&i)
	fmt.Println(32)
	show(b)
}

func testInt127() {
	i := 127
	b, _ := msgpack.Marshal(&i)
	fmt.Println(127)
	show(b)
}

func testInt128() {
	i := 128
	b, _ := msgpack.Marshal(&i)
	fmt.Println(128)
	show(b)
}

func testInt255() {
	i := 255
	b, _ := msgpack.Marshal(&i)
	fmt.Println(255)
	show(b)
}

func testIntN31() {
	i := -31
	b, _ := msgpack.Marshal(&i)
	fmt.Println(-31)
	show(b)
}

func testIntN32() {
	i := -32
	b, _ := msgpack.Marshal(&i)
	fmt.Println(-32)
	show(b)
}

func testIntN33() {
	i := -33
	b, _ := msgpack.Marshal(&i)
	fmt.Println(-33)
	show(b)
}

func main() {
	testBooleanTrue()
	testBooleanFalse()

	testInt1()
	testInt10()
	testInt31()
	testInt32()
	testInt127()
	testInt128()
	testInt255()

	testIntN31()
	testIntN32()
	testIntN33()
}
