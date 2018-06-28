package rand

/*
#include <stdlib.h>
*/
import "C"

//func Random() int {
//	return int(C.rand())
//}

func Random() int {
	var r C.int = C.rand()
	return int(r)
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
