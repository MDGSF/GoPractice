/*
修改以下代码，使避免concurrent map writes错误
package main

import (
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))
}
*/

package main

import (
	"math/rand"
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
			m[rand.Int()] = rand.Int()
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}
