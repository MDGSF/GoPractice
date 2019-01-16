package main

import (
	"fmt"
	"os"
	"sort"
)

type TOneResult struct {
	Arr       []int
	Num       int
	MinResult int
}

func main() {
	fmt.Println("vim-go")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for num := 0; num < 20; num++ {
		a1 := make([]int, len(a))
		a2 := make([]int, len(a))
		copy(a1, a)
		copy(a2, a)
		ret1 := force(a1, num)
		ret2 := my(a2, num)
		if ret1 != ret2 {
			fmt.Printf("ret1 = %v, ret2 = %v", ret1, ret2)
			os.Exit(0)
		}
	}
}

func my(a []int, num int) int {
	sort.Ints(a)
	length := len(a)
	maxNum := a[length-1]
	minNum := a[0]
	midNum := (maxNum + minNum) / 2

	if num >= midNum {
		return maxNum - minNum
	}

	for k := range a {
		if a[k] <= midNum {
			a[k] += num
		} else {
			a[k] -= num
		}
	}

	sort.Ints(a)
	return a[length-1] - a[0]
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

	//fmt.Println("number of results = ", len(results))

	first := true
	var min *TOneResult
	for k := range results {
		//fmt.Println("k =", k)
		//fmt.Println("Arr =", v.Arr)
		//fmt.Println("Num =", v.Num)
		//fmt.Println("min =", v.MinResult)
		if first {
			min = results[k]
			first = false
		} else {
			if min.MinResult > results[k].MinResult {
				min = results[k]
			}
		}
	}

	fmt.Println(a, num, min)

	return min.MinResult
}
