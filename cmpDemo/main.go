package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	type T struct {
		I int
	}

	x := []*T{{1}, {2}, {3}}
	y := []*T{{1}, {2}, {4}}

	fmt.Println(cmp.Equal(x, y))

	diff := cmp.Diff(x, y)
	fmt.Printf(diff)
}
