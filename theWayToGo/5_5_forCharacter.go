package main

func main() {
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}

	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
