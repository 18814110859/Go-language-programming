package fib

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(b, 30) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//func BenchmarkFib100(b *testing.B) {benchmarkFib(b, 100)}
//func BenchmarkFib1000(b *testing.B) {benchmarkFib(b, 1000)}
//func BenchmarkFib10000(b *testing.B) {benchmarkFib(b, 10000)}
//func BenchmarkFib100000(b *testing.B) {benchmarkFib(b, 100000)}
//func BenchmarkFib1000000(b *testing.B) {benchmarkFib(b, 1000000)}
//func BenchmarkFib10000000(b *testing.B) {benchmarkFib(b, 10000000)}
//func BenchmarkFib100000000(b *testing.B) {benchmarkFib(b, 100000000)}
