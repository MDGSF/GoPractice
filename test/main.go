package main

import "fmt"

type A struct {
	Field1 int
	Field2 string
}

func main() {
	a := &A{Field1: 5, Field2: "BB"}
	var b interface{}
	b = a

	a2 := b.(*A)
	a2.Field1 = 123

	fmt.Println(a)
	fmt.Println(a2)

	b.(*A).Field1 = 321
	fmt.Println(a)
	fmt.Println(a2)
}
