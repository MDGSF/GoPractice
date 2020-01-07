package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	test(&array)
	fmt.Println(array)
}

func test(array *[]int) {
	*array = append(*array, 6)
}
