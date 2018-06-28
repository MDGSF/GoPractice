package main

/*
#cgo CFLAGS : -I../cppLibSource
#cgo LDFLAGS: -L../cppLib/linux/dynamicLib -ladd

#include "add.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Printf("add(1, 2) = %v\n", C.add(1, 2))
}
