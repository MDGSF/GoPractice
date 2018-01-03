/*
修改以下代码，使map中值对象的字段可修改
package main

type S struct {
	name string
}

func main() {
	m := map[string]S{"x": S{"one"}}
	m["x"].name = "two" //cannot assign to struct field m["x"].name in map
}
*/

package main

type S struct {
	name string
}

func main() {
	m := map[string]*S{"x": &S{"one"}} //把map的value改为指针类型的，就可以修改了。
	m["x"].name = "two"
}
