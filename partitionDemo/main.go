package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func sort(a []int) {
	if len(a) < 2 {
		return
	}
	i := partition(a)
	sort(a[:i])
	sort(a[i+1:])
}

func TopK(a []int, k int) []int {
	if k <= 0 {
		return nil
	}

	if k > len(a) {
		return a
	}

	sa := a
	sk := k
	for {
		i := partition(a)
		if k-1 == i {
			return sa[:sk]
		} else if k-1 < i {
			a = a[:i]
		} else {
			a = a[i+1:]
			k -= (i + 1)
		}
	}
}

func partition(a []int) int {
	if len(a) < 2 {
		return 0
	}
	start := 0
	end := len(a) - 2
	lastElem := a[len(a)-1]
	p := -1
	for i := start; i <= end; i++ {
		if a[i] <= lastElem {
			a[p+1], a[i] = a[i], a[p+1]
			p++
		}
	}
	a[p+1], a[len(a)-1] = a[len(a)-1], a[p+1]
	return p + 1
}
