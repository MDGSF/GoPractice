package main

import (
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type TStudent struct {
	Name string `json:"jsonname" msgpack:"msgpackname"`
}

func main() {
	s := &TStudent{}
	s.Name = "HuangJian"

	jsonByte, _ := json.Marshal(s)
	fmt.Println(string(jsonByte))

	msgpackByte, _ := msgpack.Marshal(s)
	fmt.Println(string(msgpackByte))

	s1 := &TStudent{}
	json.Unmarshal(jsonByte, s1)
	fmt.Println(s1)

	s2 := &TStudent{}
	msgpack.Unmarshal(msgpackByte, s2)
	fmt.Println(s2)
}
