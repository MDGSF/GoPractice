package main

import "fmt"

func main() {
	s1 := `
	hello
	`
	s2 := `hello ` + "`" + `world` + "`" + ``

	fmt.Println(s1)
	fmt.Println(s2)
}
