package main

import "fmt"

func main() {
	fmt.Println("fib(0) = ", fib(0))
	fmt.Println("fib(1) = ", fib(1))
	fmt.Println("fib(2) = ", fib(2))
	fmt.Println("fib(3) = ", fib(3))
	fmt.Println("fib(4) = ", fib(4))
	fmt.Println("fib(5) = ", fib(5))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	i, j := 0, 1
	for idx := 2; idx <= n; idx++ {
		i, j = j, i+j
	}
	return j
}

func fib1(n int, m map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := m[n]; !ok {
		m[n] = fib1(n-1, m) + fib1(n-2, m)
	}
	return m[n]
}
