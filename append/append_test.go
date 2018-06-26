package main

import (
	"testing"
)

const N = 100000

func BenchmarkAppendZeroCap(b *testing.B) {
	list := make([]int, 0, 1)

	for i := 1; i < N; i++ {
		list = append(list, i)
	}
}

func BenchmarkAppendSizedCap(b *testing.B) {
	list := make([]int, 0, N)

	for i := 1; i < N; i++ {
		list = append(list, i)
	}
}
