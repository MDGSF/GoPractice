package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			ctx := context.WithValue(ctx, i, i)
			<-ctx.Done()

			fmt.Println("Cancelled:", i)
		}(i)
	}

	cancel()
	wg.Wait()
}
