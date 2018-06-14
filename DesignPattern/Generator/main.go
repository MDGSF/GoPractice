package main

import "fmt"

func Count(start int, end int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func main() {
	for i := range Count(1, 99) {
		fmt.Println(i)
	}
}
