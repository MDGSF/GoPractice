package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
	read()
}

func write() {
	f, _ := os.Create("test.csv")
	w := csv.NewWriter(f)
	data := [][]string{
		{"key_tip1", "中国", "China"},
		{"key_apple", "苹果", "Apple"},
	}
	w.WriteAll(data)
	w.Flush()
}

func read() {
	f, _ := os.Open("test.csv")
	r := csv.NewReader(f)
	data, _ := r.ReadAll()
	fmt.Println(data)
}
