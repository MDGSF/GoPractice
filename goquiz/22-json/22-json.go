/*
纠正下面代码的一个错误
package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	status int
}

func main() {
	var data = []byte(`{"status": 200}`)
	result := &Result{}

	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result=%+v", result) //输出: result=&{status:0}
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status int `json:"status"` //把Status的第一个字符改为大写，并加上`json:"status"`
}

func main() {
	var data = []byte(`{"status": 200}`)
	result := &Result{}

	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result=%+v", result) //输出: result=&{status:0}
}
