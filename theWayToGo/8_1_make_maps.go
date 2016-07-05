package main

import "fmt"

func main() {
	var mapLit map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}

	var mapAssigned map[string]int
	mapAssigned = mapLit
	mapAssigned["two"] = 3

	mapCreated := make(map[string]float32)
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map literal at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
}
