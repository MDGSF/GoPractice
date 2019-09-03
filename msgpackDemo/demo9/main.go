package main

import (
	"fmt"

	"github.com/MDGSF/utils/log"
	"github.com/vmihailenco/msgpack"
)

type Student struct {
	ID   string `msgpack:"id"`
	Name string `msgpack:"name"`
}

type LibFlowData struct {
	Key  string      `msgpack:"key"`
	Data interface{} `msgpack:"data"`
}

func main() {

	var err error

	s1 := &Student{
		ID:   "1",
		Name: "No1 aaaaaaaa",
	}

	libdata := &LibFlowData{
		Key:  "test1",
		Data: s1,
	}

	b, err := msgpack.Marshal(libdata)
	if err != nil {
		log.Error("%v", err)
		return
	}
	log.Info("b = %v", len(b))

	var s2 Student
	outlib := &LibFlowData{Data: &s2}
	err = msgpack.Unmarshal(b, outlib)
	if err != nil {
		log.Error("%v", err)
		return
	}

	fmt.Println(outlib)
	fmt.Println(s2)
}
