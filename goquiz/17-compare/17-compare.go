/*
解释以下代码打印的结果为什么是false,并修改A行使打印x和y的值是否相等的结果
package main

import (
	"fmt"
)

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(x == y) //A 这里会输出false
}
*/

package main

import (
	"fmt"
)

type S struct {
	a, b, c string
}

func main() {
	x := S{"a", "b", "c"}
	y := S{"a", "b", "c"}
	fmt.Println(x == y) //A 这里会输出true
}
