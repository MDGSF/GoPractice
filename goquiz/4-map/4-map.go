/*
在以下AB两行修改代码，用正确的用法存取map数据
package main

func main() {
	var m map[string]int //A
	m["a"] = 1
	if v := m["b"]; v != nil { //B
		println(v)
	}
}
*/

package main

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	if v, ok := m["b"]; ok {
		println(v)
	}
}
