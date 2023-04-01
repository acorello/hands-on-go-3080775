// testing/benchmarks/begin/main_test.go
package main

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2, 3, 4, 5)
	}
}

func BenchmarkSumInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumInt(1, 2, 3, 4, 5)
	}
}

func BenchmarkSumAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAny([]interface{}{1, 2, 3, 4, 5})
	}
}
