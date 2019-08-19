package main

import "fmt"

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	h.data = make([]int, 0)
	return h
}

func NewMinHeapWithArray(data []int) *MinHeap {
	h := &MinHeap{}
	h.data = make([]int, len(data))
	copy(h.data, data)
	h.heapify()
	return h
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Push(x int) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *MinHeap) Pop() int {
	n := len(h.data) - 1
	h.swap(0, n)
	h.down(0, n)
	v := 0
	h.data, v = h.data[:n], h.data[n]
	return v
}

func (h *MinHeap) heapify() {
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *MinHeap) up(i int) {
	for {
		p := h.parent(i)
		if p == i {
			break
		}

		if !h.less(i, p) {
			break
		}

		h.swap(i, p)
		i = p
	}
}

func (h *MinHeap) down(i0, n int) bool {
	i := i0
	for {
		l := h.leftChild(i)
		if l >= n {
			break
		}

		m := l

		r := h.rightChild(i)
		if r < n && h.less(r, l) {
			m = r
		}

		if !h.less(m, i) {
			break
		}

		h.swap(i, m)
		i = m
	}
	return i > i0
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func main() {
	data := []int{1, 2, 3, 4, 5, 9, 8, 7, 6}
	minh := NewMinHeapWithArray(data)
	for minh.Len() > 0 {
		fmt.Println(minh.Pop())
	}

	for i := 20; i > 10; i-- {
		minh.Push(i)
	}
	for minh.Len() > 0 {
		fmt.Printf("%v ", minh.Pop())
	}
	fmt.Println()
}
