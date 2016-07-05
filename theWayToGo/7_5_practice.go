package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}

	for i, item := range items {
		fmt.Printf("%d = %d\n", i, item)
	}
	fmt.Println()

	for ix := range items {
		items[ix] *= 2
	}

	for i, item := range items {
		fmt.Printf("%d = %d\n", i, item)
	}
}
