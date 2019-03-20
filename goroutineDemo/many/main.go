package main

import "fmt"

func main() {

	ch := make(chan struct{})
	i := 0

	for {
		go func() {
			ch <- struct{}{}
		}()
		i++
		if i%10000 == 0 {
			fmt.Println(i)
		}
		if i > 6000000 {
			break
		}
	}
}
