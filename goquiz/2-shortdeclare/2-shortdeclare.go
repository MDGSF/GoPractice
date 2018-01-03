package main

import "fmt"

func f() (int, int) {
	return 0, 1
}

func test1() {
	// var x int
	// x, _ := f() //no new variables on left side of :=
}

func test2() {
	var x int
	x, _ = f()
	fmt.Println(x)
}

func test3() {
	var x int
	x, y := f()
	fmt.Println(x, y)
}

func test4() {
	// var x int
	// x, y = f() //undefined: y
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
