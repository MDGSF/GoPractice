package main

import (
	"fmt"
)

const IntMax = int32(^uint32(0) >> 1)
const IntMin = ^IntMax

func main() {
	fmt.Println("IntMax =", IntMax)
	fmt.Println("IntMin =", IntMin)
}
