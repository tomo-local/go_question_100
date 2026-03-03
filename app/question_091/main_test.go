package main

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}
