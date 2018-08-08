package main

import "fmt"

func main() {
	c1 := make(chan int)
	test(c1, c1)
}

func test(c1 chan int, c2 chan int) {
	if c1 == c2 {
		fmt.Println("c1==c2")
	}
}
