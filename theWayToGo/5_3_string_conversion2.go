package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	//var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exit with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
}
