package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func st(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	s1 := arr[len(arr)-1] + st(arr[:len(arr)-2])
	s2 := st(arr[:len(arr)-1])
	return max(s1, s2)
}

func st2(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	fn := make([]int, len(arr))
	fn[0] = arr[0]
	fn[1] = max(arr[0], arr[1])
	return st2Inner(arr, len(arr), fn)
}

func st2Inner(arr []int, n int, fn []int) int {
	if n <= 0 {
		return 0
	}
	if fn[n-1] != 0 {
		return fn[n-1]
	}
	s1 := arr[n-1] + st2Inner(arr, n-2, fn)
	s2 := st2Inner(arr, n-1, fn)
	fn[n-1] = max(s1, s2)
	return fn[n-1]
}

func st3(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	fn := make([]int, len(arr))
	fn[0] = arr[0]
	fn[1] = max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		fn[i] = max(arr[i]+fn[i-2], fn[i-1])
	}
	return fn[len(arr)-1]
}

func st4(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	p2 := arr[0]
	p1 := max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		cur := max(arr[i]+p2, p1)
		p2 = p1
		p1 = cur
	}
	return p1
}

func main() {
	fmt.Println("hello world")
}
