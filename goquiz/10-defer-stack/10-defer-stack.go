/*
纠正下面代码的一个错误
package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
*/

package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	if err != nil {
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
