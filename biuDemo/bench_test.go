package main

import (
	"testing"

	"github.com/imroc/biu"
)

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 0)
		b := byte('3')
		appendBinaryString(bs, b)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 0)
		b := byte('3')
		appendBinaryString2(bs, b)
	}
}

func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 8)
		b := byte('3')
		appendBinaryString3(bs, b)
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biu.ToBinaryString(i)
	}
}
