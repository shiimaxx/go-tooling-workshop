package main

import "testing"

var s string

func Benchmark_isGopher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, _ = isGopher("yoshima@golang.org")
	}
}
