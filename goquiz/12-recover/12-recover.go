package main

import "log"

func f() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}

/*
输出:
2018/01/03 16:21:50 recover:1
*/
