package main

import (
	"container/list"
	"fmt"
)

/*
Node 节点

4 2
1
2 3
4 5 6
7 8 9 10

1 1 1
2 2 3
3 3 6
4 3 7
5 4 12
6
*/
type Node struct {
	i       int
	j       int
	num     int
	pathnum int
	pathlen int
}

var arr [2000][2000]*Node

func main() {
	var x int
	var y int
	fmt.Scanf("%d %d", &x, &y)
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			var num int
			fmt.Scanf("%d", &num)
			arr[i][j] = &Node{}
			arr[i][j].i = i
			arr[i][j].j = j
			arr[i][j].num = num
		}
	}

	// x = 4
	// y = 2
	// arr[0][0] = &Node{i: 0, j: 0, num: 1, pathnum: 1}

	// arr[1][0] = &Node{i: 1, j: 0, num: 2}
	// arr[1][1] = &Node{i: 1, j: 1, num: 3}

	// arr[2][0] = &Node{i: 2, j: 0, num: 4}
	// arr[2][1] = &Node{i: 2, j: 1, num: 5}
	// arr[2][2] = &Node{i: 2, j: 2, num: 6}

	// arr[3][0] = &Node{i: 3, j: 0, num: 7}
	// arr[3][1] = &Node{i: 3, j: 1, num: 8}
	// arr[3][2] = &Node{i: 3, j: 2, num: 9}
	// arr[3][3] = &Node{i: 3, j: 3, num: 10}

	//Show()

	ret := f(x, y)
	fmt.Printf("%v", ret)
}

func f(x, y int) int {
	l := &list.List{}
	l.Init()
	l.PushBack(arr[0][0])
	arr[0][0].pathnum = 1
	arr[0][0].pathlen = arr[0][0].num

	for l.Len() > 0 {
		newl := &list.List{}
		newl.Init()

		for l.Len() > 0 {
			e := l.Front()
			n := e.Value.(*Node)
			process(n, arr[n.i][n.j+1], newl)
			process(n, arr[n.i+1][n.j], newl)
			l.Remove(e)
		}

		if newl.Len() > 0 {
			firstE := newl.Front()
			firstV := firstE.Value.(*Node)
			if firstV.pathnum == y {

				minVal := firstV.pathlen
				newl.Remove(firstE)
				for newl.Len() > 0 {
					curE := newl.Front()
					curV := curE.Value.(*Node)
					if curV.pathlen < minVal {
						minVal = curV.pathlen
					}
					newl.Remove(curE)
				}
				return minVal

			} else {
				l.PushBackList(newl)
			}
		}
	}

	return 0
}

func process(prev *Node, cur *Node, newl *list.List) {
	if cur == nil {
		return
	}

	if cur.pathnum == 0 || cur.pathlen > prev.pathlen+cur.num {
		cur.pathnum = prev.pathnum + 1
		cur.pathlen = prev.pathlen + cur.num
		newl.PushBack(cur)
	}
}

func Show() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if arr[i][j] == nil {
				continue
			}
			fmt.Printf("%d ", arr[i][j].num)
		}
		fmt.Println()
	}
	fmt.Println()
}
