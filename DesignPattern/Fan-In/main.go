package main

import (
	"fmt"
	"sync"
)

// Merge different channels in one channel.
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	c := Merge(c1, c2, c3)

	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
			c2 <- i
			c3 <- i
		}
		close(c1)
		close(c2)
		close(c3)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
