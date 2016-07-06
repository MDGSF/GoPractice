package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	fmt.Printf("im: %v\n", im)
	fmt.Printf("in: %v\n", in)
	fmt.Printf("io: %v\n", io)
	fmt.Printf("ip: %v\n\n", ip)

	//one := big.NewInt(1)
	//two := big.NewInt(2)
	//three := big.NewInt(3)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("im: %v\n", im)
	fmt.Printf("in: %v\n", in)
	fmt.Printf("io: %v\n", io)
	fmt.Printf("ip: %v\n", ip)

	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}
