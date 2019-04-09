package main

import (
	"bytes"
	"fmt"

	"github.com/bradleyjkemp/memviz"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	out := &bytes.Buffer{}
	memviz.Map(out, &a)
	fmt.Println(out.String())
}
