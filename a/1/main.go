package main

import (
	"fmt"
	"sort"
)

type TOneResult struct {
	Arr       []int
	Num       int
	MinResult int
}

func main() {
	fmt.Println("vim-go")

	a := []int{1, 2, 3, 4, 5, 6}
	for num := 0; num < 20; num++ {
		force(a, num)
	}
}

func force(a []int, num int) int {

	results := make([]*TOneResult, 0)

	var t func(a []int, num int, idx int)
	t = func(a []int, num int, idx int) {
		if idx < len(a) {
			a[idx] += num
			t(a, num, idx+1)
			a[idx] -= num

			a[idx] -= num
			t(a, num, idx+1)
			a[idx] += num
			return
		}

		oneResult := &TOneResult{}
		oneResult.Arr = make([]int, len(a))
		copy(oneResult.Arr, a)
		oneResult.Num = num
		sort.Ints(oneResult.Arr)
		oneResult.MinResult = oneResult.Arr[len(oneResult.Arr)-1] - oneResult.Arr[0]
		results = append(results, oneResult)
	}

	t(a, num, 0)

	fmt.Println("number of results = ", len(results))
	for k, v := range results {
		fmt.Println("k =", k)
		fmt.Println("Arr =", v.Arr)
		fmt.Println("Num =", v.Num)
		fmt.Println("min =", v.MinResult)
	}

	return 0
}
