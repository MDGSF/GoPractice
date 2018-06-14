package main

import (
	"fmt"

	"github.com/MDGSF/GoPractice/DesignPattern/Singleton"
)

func main() {
	s := Singleton.New()
	s["this"] = "that"

	s2 := Singleton.New()
	fmt.Println("This is", s2["this"])
}
