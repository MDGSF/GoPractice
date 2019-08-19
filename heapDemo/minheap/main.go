package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := NewMyMinHeap(3)
	for i := 11; i < 20; i++ {
		h.PushOneElement(i)
		fmt.Println(h)
	}
	for i := 1; i < 10; i++ {
		h.PushOneElement(i)
	}
	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Printf("%v ", h.PopOneElement())
	}
	fmt.Println()
}

/*
MyMinHeap 最小堆，限制堆的大小
用来计算一个数据流中最大的几个数
*/
type MyMinHeap struct {
	elems   []int
	maxSize int
}

// NewMyMinHeap create new my heap
func NewMyMinHeap(maxSize int) *MyMinHeap {
	h := &MyMinHeap{}
	h.elems = make([]int, 0)
	h.maxSize = maxSize
	return h
}

func (h *MyMinHeap) Len() int {
	return len(h.elems)
}

func (h *MyMinHeap) Less(i, j int) bool {
	return h.elems[i] < h.elems[j]
}

func (h *MyMinHeap) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

// Push add x as element Len()
func (h *MyMinHeap) Push(x interface{}) {
	h.elems = append(h.elems, x.(int))
}

// Pop remove and return element Len() - 1.
func (h *MyMinHeap) Pop() (v interface{}) {
	h.elems, v = h.elems[:h.Len()-1], h.elems[h.Len()-1]
	return
}

// PushOneElement push
func (h *MyMinHeap) PushOneElement(x interface{}) {
	if len(h.elems) >= h.maxSize {
		if h.elems[0] >= x.(int) {
			return
		}
		heap.Pop(h)
	}
	heap.Push(h, x)
}

// PopOneElement pop
func (h *MyMinHeap) PopOneElement() interface{} {
	return heap.Pop(h)
}
