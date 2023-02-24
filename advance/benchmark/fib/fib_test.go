package fib

import (
	"testing"
)

func FibRec(n int) int {  // 递归
	if n < 2 {
		return n
	}

	return FibRec(n-1) + FibRec(n-2)
}


func BenchmarkFibRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRec(10)
	}
}


func FibLoop(n int) int {  // 循环
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}


func BenchmarkFibLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibLoop(10)
	}
}


// go test -v -bench ^BenchmarkFibRec$ -benchmem -run=^$