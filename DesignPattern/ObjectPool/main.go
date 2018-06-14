package main

import (
	"fmt"

	"github.com/MDGSF/GoPractice/DesignPattern/ObjectPool/pool"
)

func main() {
	p := pool.New(2)

	for {
		select {
		case obj := <-p:
			//obj.Do()
			fmt.Println("get one object = ", obj)
			p <- obj
		default:
			fmt.Println("has no object")
			return
		}
	}
}
