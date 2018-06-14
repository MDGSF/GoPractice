package main

import (
	"fmt"

	"github.com/MDGSF/GoPractice/DesignPattern/StrategyPattern"
)

func main() {
	add := StrategyPattern.Operation{StrategyPattern.Addition{}}
	r := add.Operate(3, 5)
	fmt.Println(r)
}
