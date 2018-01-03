package main

func main() {

Outter:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break Outter //跳出两层循环，不会再次进入for循环。
		}
		println()
	}

}
