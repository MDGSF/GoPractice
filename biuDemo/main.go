package main

import (
	"fmt"

	"github.com/imroc/biu"
)

func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, '0')
		default:
			bs = append(bs, '1')
		}
		b <<= 1
	}
	return bs
}

func appendBinaryString2(bs []byte, b byte) []byte {
	for i := 7; i >= 0; i-- {
		if b&(1<<i) == 0 {
			bs = append(bs, '0')
		} else {
			bs = append(bs, '1')
		}
	}
	return bs
}

func appendBinaryString3(bs []byte, b byte) []byte {
	for i := 7; i >= 0; i-- {
		if b&0x01 == 0 {
			bs[i] = '0'
		} else {
			bs[i] = '1'
		}
		b >>= 1
	}
	return bs
}

func main() {
	fmt.Println(biu.ToBinaryString(3))

	bs := make([]byte, 0)
	b := byte('3')
	out := appendBinaryString(bs, b)
	fmt.Println(out)

	bs2 := make([]byte, 0)
	out2 := appendBinaryString2(bs2, b)
	fmt.Println(out2)

	bs3 := make([]byte, 8)
	out3 := appendBinaryString3(bs3, b)
	fmt.Println(out3)
}
