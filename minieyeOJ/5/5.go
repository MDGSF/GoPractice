package main

import (
	"container/list"
	"fmt"
)

type Pair struct {
	k int
	v int
}

var arr [30]int

func main() {
	var x int
	var y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%v", f(x, y))
}

func showArr() {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for k := range arr {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
}

func f(x, y int) int {

	for k := range arr {
		arr[k] = -1
	}
	arr[x] = 0

	l := &list.List{}
	l.Init()
	l.PushBack(Pair{k: x, v: 0})

	for l.Len() > 0 {
		p := l.Front().Value.(Pair)

		if v, end := process(l, y, p.k+1, p.v+1); end {
			return v
		}
		if v, end := process(l, y, p.k-1, p.v+1); end {
			return v
		}
		if v, end := process(l, y, p.k*2, p.v+1); end {
			return v
		}

		showArr()
	}

	return 0
}

func process(l *list.List, y int, newkey int, newval int) (int, bool) {

	if newkey < 0 || newkey > 100000 {
		return 0, false
	}

	if arr[newkey] == -1 || arr[newkey] > newval {
		arr[newkey] = newval
		l.PushBack(Pair{k: newkey, v: newval})
	}

	if newkey == y {
		return arr[newkey], true
	}
	return 0, false
}
