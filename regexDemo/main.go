package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("vim-go")

	r, _ := regexp.Compile(`^(.+?)\s+(.+)$`)
	all := r.FindAllString("127.0.0.1:12580  desekldjsflkasjf:fdasf  fkdasjflksajf", -1)
	fmt.Println(len(all), all)
}
