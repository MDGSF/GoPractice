/*
https://github.com/google/re2/wiki/Syntax

^	at beginning of text or line (m=true)
$	at end of text (like \z not \Z) or line (m=true)

.         any character, possibly including newline (s=true)
[xyz]     character class
[^xyz]    negated character class

xy        x followed by y
x|y       x or y (prefer x)

x*        zero or more x, prefer more
x+        one or more x, prefer more
x?        zero or one x, prefer one
x{n,m}	  n or n+1 or ... or m x, prefer more
x{n,}	  n or more x, prefer more
x{n}	  exactly n x
x*?	      zero or more x, prefer fewer
x+?	      one or more x, prefer fewer
x??	      zero or one x, prefer zero
x{n,m}?	  n or n+1 or ... or m x, prefer fewer
x{n,}?	  n or more x, prefer fewer
x{n}?	  exactly n x

\d	digits (≡ [0-9])
\D	not digits (≡ [^0-9])
\s	whitespace (≡ [\t\n\f\r ])
\S	not whitespace (≡ [^\t\n\f\r ])
\w	word characters (≡ [0-9A-Za-z_])
\W	not word characters (≡ [^0-9A-Za-z_])
*/
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	fmt.Println("vim-go")

	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}

func test8() {
	// ip address
	r := regexp.MustCompile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`)
	all := r.FindAllString("my server ip addr 127.0.0.1, and port is 22", -1)
	fmt.Println(len(all), all)
}

func test7() {
	// 把多个空格替换为一个空格
	r := regexp.MustCompile(` +`)
	all := r.ReplaceAllString("Welcome   to   Beijing", " ")
	fmt.Println(all)
	// Welcome to Beijing
}

func test6() {
	// 把空格替换为@
	r := regexp.MustCompile(" ")
	all := r.ReplaceAllString("Welcome to Beijing", "@")
	fmt.Println(all)
	// Welcome@to@Beijing
}

func test5() {
	// 查找小写字符和数字
	r, err := regexp.Compile(`[a-z]+|[0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	all := r.FindAllString("AA12bb34CC56dd78", -1)
	fmt.Println(len(all), all)
	// 6 [12 bb 34 56 dd 78]
}

func test4() {
	// 查找单词
	r, _ := regexp.Compile(`\w+`)
	all := r.FindAllString("this is a beautiful day.", -1)
	fmt.Println(len(all), all)
	//5 [this is a beautiful day]
}

func test3() {
	// 把相连的字符和数字分开
	r, _ := regexp.Compile(`\D+|\d+`)
	all := r.FindAllString("aa12bb34cc56", -1)
	fmt.Println(len(all), all)
	// 6 [aa 12 bb 34 cc 56]
}

func test2() {
	r, _ := regexp.Compile(`^\s*\"(.+?)\"\s*$`)
	fmt.Println(r.MatchString("  \"Huang Jian\"  "))
	// true
}

func test1() {
	r, _ := regexp.Compile(`^(.+?)\s+(.+)$`)
	fmt.Println(r.MatchString("127.0.0.1:12580  desekldjsflkasjf:fdasf  fkdasjflksajf"))
	// true
}
