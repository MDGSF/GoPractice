package workerpool

import (
	"fmt"
	"time"
)

type Job struct {
	Payload Payload
}

type Payload int

func (p Payload) Do() (err error) {
	if int(p)%100000 == 0 {
		fmt.Println("I am working do", int(p))
	}

	// suppose need to use 20 millisecond to process.
	time.Sleep(20 * time.Millisecond)

	err = nil
	return
}
