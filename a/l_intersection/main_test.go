package main

import "testing"

func CreateList(arr []int) (*ListNode, *ListNode) {
	var head *ListNode
	var tail *ListNode
	for k := range arr {
		curNode := &ListNode{Value: arr[k]}
		if k == 0 {
			head = curNode
			tail = curNode
		} else {
			tail.next = curNode
			tail = curNode
		}
	}
	return head, tail
}

func Test1(t *testing.T) {
	result := FindIntersection(nil, nil)
	if result != nil {
		t.Fatal("invalid result =", result)
	}
}

func Test2(t *testing.T) {
	t1, _ := CreateList([]int{1, 2, 3, 4})
	result := FindIntersection(t1, nil)
	if result != nil {
		t.Fatal("invalid result =", result)
	}
}

func Test3(t *testing.T) {
	t1, _ := CreateList([]int{1, 2, 3, 4})
	result := FindIntersection(nil, t1)
	if result != nil {
		t.Fatal("invalid result =", result)
	}
}

func Test4(t *testing.T) {
	t1, _ := CreateList([]int{1, 2, 3, 4})
	t2, t2Tail := CreateList([]int{7, 8, 9})
	t2Tail.next = t1
	result := FindIntersection(t1, t2)
	if result == nil {
		t.Fatal("invalid result =", result)
	}
	if result.Value != 1 {
		t.Fatal("invalid result =", result.Value)
	}
}

func Test5(t *testing.T) {
	t1, _ := CreateList([]int{1, 2, 3, 4})
	t2, t2Tail := CreateList([]int{7, 8, 9})
	t2Tail.next = t1.next
	result := FindIntersection(t1, t2)
	if result == nil {
		t.Fatal("invalid result =", result)
	}
	if result.Value != 2 {
		t.Fatal("invalid result =", result.Value)
	}
}

func Test6(t *testing.T) {
	t1, _ := CreateList([]int{1, 2, 3, 4})
	t2, _ := CreateList([]int{7, 8, 9})
	result := FindIntersection(t1, t2)
	if result != nil {
		t.Fatal("invalid result =", result)
	}
}
