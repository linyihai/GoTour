package gotest

import (
	"testing"
)

func BenchmarkRange(b *testing.B) {
	vs := make([]int, 100000)
	null := 0
	for i := 0; i < b.N; i++ {
		for _, v := range vs {
			null = v + null
		}
	}
}

func BenchmarkFor(b *testing.B) {
	vs := make([]int, 10000)
	null := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(vs); j++ {
			null = null + j
		}
	}
}

func BenchmarkLoopSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		total := 0
		for j := 0; j <= maxLoop; j++ {
			total += j
		}
	}
}

var maxLoop = 1000000

func recursionSum(sum, iterNum *int) {
	if *iterNum > maxLoop {
		return
	}

	*sum += *iterNum
	*iterNum++
	recursionSum(sum, iterNum)
}

func BenchmarkLoopRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		total := 0
		iterNum := 0
		recursionSum(&total, &iterNum)
	}
}
