package main

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/slene/blackfriday"
)

func main() {
	fmt.Println("vim-go")

	input, _ := ioutil.ReadFile("2016-05-13-c-bit-byteDemo.md")
	output := blackfriday.MarkdownBasic(input)

	//fmt.Println(string(output))

	output2 := beego.HTML2str(string(output))
	fmt.Println(output2)
}
