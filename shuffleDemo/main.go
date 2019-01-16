package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a)

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(a)-1, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	fmt.Println(a)
}
