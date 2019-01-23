package main

func main() {
	test()
}

func test() {
	println("test start")
	defer println("test end")

	for i := 0; i < 5; i++ {
		println("i =", i)
		defer println("defer i =", i)
	}

	println("end of loop")
}
