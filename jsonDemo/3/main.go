package main

import (
	"encoding/json"
	"fmt"
)

type TSubNode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TSubNode2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TSubNode3 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TNode struct {
	Sub  TSubNode  `json:"subnode,inline,omitempty"`
	Sub2 TSubNode2 `json:"subnode2"`
	Sub3 TSubNode3 `json:",inline"`
	ID   int       `json:"id"`
}

func main() {
	t1 := &TNode{
		ID: 123,
	}

	data1, _ := json.Marshal(t1)
	fmt.Println(string(data1))

	t2 := &TNode{}
	json.Unmarshal(data1, t2)

	fmt.Println(t2)
}
