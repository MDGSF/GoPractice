package main

import "fmt"

func main() {
	s := "good bye"
	fmt.Printf("Here is the string s: %s\n", s)

	var p *string = &s
	*p = "ciao"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
}
