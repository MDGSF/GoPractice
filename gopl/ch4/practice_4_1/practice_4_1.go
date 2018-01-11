package main

import (
	"crypto/sha256"
	"fmt"
)

func getOneBit(num byte, i uint) int {
	mask := uint(1) << i
	if int(uint(num)&mask) > 0 {
		return 1
	}
	return 0
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T %x\n", c1, c1)
	fmt.Printf("%T %x\n", c2, c2)

	count := 0
	for k, v := range c1 {
		for i := 0; i < 8; i++ {
			if getOneBit(v, uint(i)) != getOneBit(c2[k], uint(i)) {
				count++
			}
		}
	}
	fmt.Println("count = ", count)
}
