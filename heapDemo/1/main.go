package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	//heap.Interface

	h := new(myHeap)
	for i := 11; i < 20; i++ {
		h.Push(i)
	}
	for i := 1; i < 10; i++ {
		h.Push(i)
	}
	fmt.Println(h)

	heap.Init(h)
	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(h))
	}
	fmt.Println()
}

type myHeap []int

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// add x as element Len()
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// remove and return element Len() - 1.
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
