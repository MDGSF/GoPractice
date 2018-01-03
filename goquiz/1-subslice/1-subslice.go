package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)                                             //1 12 13
	fmt.Println(ss)                                            //12 13
	fmt.Println("&s[0], &s[1], &s[2] = ", &s[0], &s[1], &s[2]) //0xc042012440 0xc042012448 0xc042012450
	fmt.Println("&ss[0], &ss[1] = ", &ss[0], &ss[1])           //0xc042012448 0xc042012450

	ss = append(ss, 4) //当执行了append之后，ss的内存地址就改变了，被重新拷贝了一份。
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)                                                   //1 12 13
	fmt.Println(ss)                                                  //22 23 14
	fmt.Println("&s[0], &s[1], &s[2] = ", &s[0], &s[1], &s[2])       //0xc042012440 0xc042012448 0xc042012450
	fmt.Println("&ss[0], &ss[1], &ss[2] = ", &ss[0], &ss[1], &ss[2]) //0xc0420124a0 0xc0420124a8 0xc0420124b0
}
