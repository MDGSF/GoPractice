package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
	L     *Node
	R     *Node
}

func create() *Node {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}

	n2.L = n1
	n2.R = n3

	n4.L = n2
	n4.R = n5

	return n4
}

func preOrder(n *Node) {
	if n == nil {
		return
	}

	preOrder(n.L)
	fmt.Printf("%v ", n.Value)
	preOrder(n.R)
}

func preOrder2List(n *Node, l *list.List) {
	if n == nil {
		return
	}

	preOrder2List(n.L, l)
	l.PushBack(n)
	preOrder2List(n.R, l)
}

func main() {
	fmt.Println("vim-go")

	n := create()
	preOrder(n)
	fmt.Println()

	l := &list.List{}
	l.Init()
	preOrder2List(n, l)

	fmt.Println("list len =", l.Len())
	curE := l.Front()
	for curE != nil {
		cur := curE.Value.(*Node)
		fmt.Printf("%v ", cur.Value)

		curE = curE.Next()
	}
	fmt.Println()
}
