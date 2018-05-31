package main

import "testing"

func BenchmarkFindCores(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindCores(1000000)
	}
}
