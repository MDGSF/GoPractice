package main

import "fmt"

func main() {
	result := make(map[string]interface{})
	result["key"] = test()
	fmt.Printf("vim-go, %v\n", result)
}

func test() interface{} {
	m := make(map[string]interface{})
	m["name"] = "huangjian"
	return m
}
