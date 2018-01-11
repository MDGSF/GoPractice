package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T %x\n", c1, c1)
	fmt.Printf("%T %x\n", c2, c2)
}
