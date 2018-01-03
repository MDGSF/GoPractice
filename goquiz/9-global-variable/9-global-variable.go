package main

import "fmt"

type G struct {
}

//var g1 //unexpected newline, expecting type

var g2 G //undefined G

var g3 = G{}

//g4 := G{} //non-declaration statement outside function body

func main() {

	fmt.Println(g2, g3)
}
