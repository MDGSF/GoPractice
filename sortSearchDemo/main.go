package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	fmt.Println("vim-go")

	data := make([]int, 0)
	data = append(data, -1)
	for i := 24; i < 100; i++ {
		data = append(data, i)
	}

	x := 23
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
	if i < len(data) && data[i] == x {
		// x is present at data[i]
		log.Println("in array", x, i, data[i])
	} else {
		// x is not present in data,
		// but i is the index where it would be inserted.
		log.Println("not in array", x, i)
	}

	ret := sort.SearchInts(data, 23)
	log.Println(ret)
}
