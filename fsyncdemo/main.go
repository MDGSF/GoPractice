package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println("vim-go")
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte("hello"))

	err = syscall.Fsync(int(f.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}

	f.Close()
}
