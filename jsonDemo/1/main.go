package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	s := &Student{}
	s.ID = 123
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}
