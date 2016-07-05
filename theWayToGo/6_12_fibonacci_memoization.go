package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]int

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= 40; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
