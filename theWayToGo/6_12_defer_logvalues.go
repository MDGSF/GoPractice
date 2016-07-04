package main

import (
	"io"
	"log"
)

func fun1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	fun1("Go")
}
