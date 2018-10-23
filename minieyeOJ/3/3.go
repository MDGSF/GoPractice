package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("%v", valid(s, "helloworld"))
}

func valid(s, sub string) bool {
	for _, c := range sub {
		index := strings.IndexRune(s, c)
		if index == -1 {
			return false
		}
		s = s[index+1:]
	}
	return true
}
