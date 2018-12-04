package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	fmt.Println("vim-go")

	p := sync.Pool{
		New: func() interface{} {
			return "huangjian"
		},
	}

	p.Put("tt1")
	p.Put("tt2")
	p.Put("tt3")
	p.Put("tt3")

	i1 := p.Get()
	i2 := p.Get()
	i3 := p.Get()
	i4 := p.Get()
	log.Println(i1)
	log.Println(i2)
	log.Println(i3)
	log.Println(i4)
}
