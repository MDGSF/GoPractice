package main

/*
#cgo CFLAGS : -I../cppLibSource
#cgo LDFLAGS: -L../cppLib/staticLib/add.a -lstdc++
#include "add.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println("add(1, 2) = %v", C.add(1, 2))
}
