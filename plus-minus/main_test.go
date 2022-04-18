package main

import "testing"

func Benchmark_plusMinus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusMinus(arr)
	}
}
