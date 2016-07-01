package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}

	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}
