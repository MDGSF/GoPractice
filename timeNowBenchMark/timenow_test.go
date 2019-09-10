package main

import (
	"testing"
	"time"
)

func Benchmark_timenow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}
