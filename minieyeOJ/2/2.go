package main

import "fmt"

func calcFiveNum(num int) int {
	fiveNum := 0
	for {
		t := num % 5
		if t != 0 {
			break
		}

		num = num / 5
		fiveNum++
		if num <= 0 {
			break
		}
	}
	return fiveNum
}

func calcZeroNum(num int) int {
	fiveNum := 0

	for i := 1; i <= num; i++ {
		fiveNum += calcFiveNum(i)
	}

	return fiveNum
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	fmt.Print(calcZeroNum(num))
}
