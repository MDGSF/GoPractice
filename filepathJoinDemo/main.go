package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(filepath.Join("a", "b"))
	fmt.Println(filepath.Join(".", "a", "b"))
	fmt.Println(filepath.Join(".", "a", "b", "c"))
	fmt.Println(filepath.Join("..", "a", "b", "c"))
}
