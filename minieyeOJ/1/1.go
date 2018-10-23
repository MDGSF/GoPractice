package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Print(add(a, b))
}
