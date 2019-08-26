package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	ch := make(chan int)

	go func() {
		for i := 1; i < 10000; i++ {
			ch <- i
			//time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	for {
		fmt.Println("main, ", <-ch)
	}
}
