/*
//解释为何以下代码打印map的值都为3，并在A行附近修改代码，使打印结果为012的序列
package main

const N = 3

func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		m[i] = &i //A
	}

	for _, v := range m {
		print(*v)
	}
}
*/

package main

const N = 3

func main() {
	m := make(map[int]int)

	for i := 0; i < N; i++ {
		m[i] = i //A
	}

	for _, v := range m {
		print(v)
	}
}
