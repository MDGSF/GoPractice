package main

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func main() {
	//println(InitType() == nil) //cannot convert nil to type S
	println(InitPointer() == nil)      //true
	println(InitEfaceType() == nil)    //false
	println(InitEfacePointer() == nil) //false
	println(InitIfaceType() == nil)    //false
	println(InitIfacePointer() == nil) //false
}
