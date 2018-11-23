package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	arrByte, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
		return
	}

	arrStr := string(arrByte)

	var root []int
	err = json.Unmarshal([]byte(arrStr), &root)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("root =", root)
	log.Printf("root = %T", root)
}

func test1() {
	fmt.Println("vim-go")

	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	arrByte, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
		return
	}

	arrStr := string(arrByte)

	var root interface{}
	err = json.Unmarshal([]byte(arrStr), &root)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("root =", root)
	log.Printf("root = %T", root)

	result, ok := root.([]interface{})
	if !ok {
		log.Println("not ok")
		return
	}
	log.Println("result =", result)

	for _, v := range result {
		log.Printf("v = %T %v %v", v, v, v.(float64))
	}
}
