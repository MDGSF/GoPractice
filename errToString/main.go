package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	err = errors.New("I'm an error")
	errStr := fmt.Sprintf("%v", err)
	fmt.Println("errStr = ", errStr)
}
