package main

import (
	"fmt"

	"github.com/MDGSF/GoPractice/cgoDemo/demo2/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println(rand.Random())
}
