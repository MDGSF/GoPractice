package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		a := rand.Perm(10)
		fmt.Println(a)
	}
}
