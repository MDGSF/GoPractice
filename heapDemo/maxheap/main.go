package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := NewMyMaxHeap(3)
	for i := 20; i >= 11; i-- {
		h.PushOneElement(i)
		fmt.Println(h)
	}
	for i := 10; i >= 1; i-- {
		h.PushOneElement(i)
		fmt.Println(h)
	}
	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Printf("%v ", h.PopOneElement())
	}
	fmt.Println()
}

/*
MyMaxHeap 最大堆，用来从数据流中找出最小的 n 个数字
*/
type MyMaxHeap struct {
	elems   []int
	maxSize int
}

// NewMyMaxHeap create new my heap
func NewMyMaxHeap(maxSize int) *MyMaxHeap {
	h := &MyMaxHeap{}
	h.elems = make([]int, 0)
	h.maxSize = maxSize
	return h
}

func (h *MyMaxHeap) Len() int {
	return len(h.elems)
}

func (h *MyMaxHeap) Less(i, j int) bool {
	return h.elems[i] >= h.elems[j]
}

func (h *MyMaxHeap) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

// Push add x as element Len()
func (h *MyMaxHeap) Push(x interface{}) {
	h.elems = append(h.elems, x.(int))
}

// Pop remove and return element Len() - 1.
func (h *MyMaxHeap) Pop() (v interface{}) {
	h.elems, v = h.elems[:h.Len()-1], h.elems[h.Len()-1]
	return
}

// PushOneElement push
func (h *MyMaxHeap) PushOneElement(x interface{}) {
	if len(h.elems) >= h.maxSize {
		if h.elems[0] <= x.(int) {
			return
		}
		heap.Pop(h)
	}
	heap.Push(h, x)
}

// PopOneElement pop
func (h *MyMaxHeap) PopOneElement() interface{} {
	return heap.Pop(h)
}
