package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B cannot use s (type S) as type *interface{} in argument to g: *interface {} is pointer to interface, not interface
	f(p) //C
	//g(p) //D cannot use p (type *S) as type *interface{} in argument to g: *interface {} is pointer to interface, not interface
}
