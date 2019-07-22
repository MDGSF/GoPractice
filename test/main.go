package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("t = ", t.Unix())
	fmt.Println("t = ", t.Nanosecond())
}
