package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("main start")
	count := 0
	err := filepath.Walk("hj2", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("main end, count =", count)
}
