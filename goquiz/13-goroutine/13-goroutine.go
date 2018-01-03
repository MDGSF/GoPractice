/*
修改以下代码，使打印的map的长度为N=10
package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}
*/

package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			m[num] = num
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}
}
