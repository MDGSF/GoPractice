package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 10
	m[1] = 999

	for k, v := range m {
		fmt.Println(k, v)
	}
}
