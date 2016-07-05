package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("arrAge %d is %d\n", i, arrAge[i])
	}
	fmt.Println()

	var arrLazy = [...]int{5, 6, 7, 8, 22}
	for i := 0; i < len(arrLazy); i++ {
		fmt.Printf("arrLazy %d is %d\n", i, arrLazy[i])
	}
	fmt.Println()

	var arrLazy2 = []int{5, 6, 7, 8, 22}
	for i := 0; i < len(arrLazy2); i++ {
		fmt.Printf("arrLazy2 %d is %d\n", i, arrLazy2[i])
	}
	fmt.Println()

	var arrKeyValue = [5]string{3: "Chirs", 4: "Ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("arrKeyValue %d is %s\n", i, arrKeyValue[i])
	}
	fmt.Println()

	var arrKeyValue2 = []string{3: "Chirs", 4: "Ron"}
	for i := 0; i < len(arrKeyValue2); i++ {
		fmt.Printf("arrKeyValue2 %d is %s\n", i, arrKeyValue2[i])
	}
	fmt.Println()
}
