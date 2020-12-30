package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("vim-go")
	a := big.NewInt(1)

	a.Add(a, big.NewInt(222))

	fmt.Println(a)
}
