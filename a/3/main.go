package main

import "fmt"

func main() {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("vim-go", r)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // num, index
	for k, v := range nums {
		o := target - v
		if k2, ok := m[o]; ok {
			return []int{k, k2}
		}
		m[v] = k
	}
	return nil
}
