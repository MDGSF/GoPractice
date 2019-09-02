package main

import "fmt"

type ListNode struct {
	Value int
	next  *ListNode
}

func Len(t *ListNode) int {
	count := 0
	for t != nil {
		count++
		t = t.next
	}
	return count
}

func FindIntersection(t1, t2 *ListNode) *ListNode {
	if t1 == nil || t2 == nil {
		return nil
	}

	len1 := Len(t1)
	len2 := Len(t2)
	shortlen := 0
	longlen := 0
	var shortlist *ListNode
	var longlist *ListNode
	if len1 > len2 {
		shortlen = len2
		longlen = len1
		shortlist = t2
		longlist = t1
	} else {
		shortlen = len1
		longlen = len2
		shortlist = t1
		longlist = t2
	}

	diff := longlen - shortlen

	for ; diff > 0; diff-- {
		longlist = longlist.next
	}

	for ; shortlen > 0; shortlen-- {
		if shortlist == longlist {
			return shortlist
		}
		shortlist = shortlist.next
		longlist = longlist.next
	}

	return nil
}

func main() {
	fmt.Println("vim-go")
}
