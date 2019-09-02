package main

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func preOrder(pRoot *Node) {
	if pRoot == nil {
		return
	}
	fmt.Printf("%v ", pRoot.Value)
	preOrder(pRoot.left)
	preOrder(pRoot.right)
}

func midOrder() {

}

func postOrder() {

}

func main() {
	fmt.Println("vim-go")
}
